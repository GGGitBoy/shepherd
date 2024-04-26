package cce

import (
	"time"

	ccev1 "github.com/cnrancher/cce-operator/pkg/apis/cce.pandaria.io/v1"
	"github.com/cnrancher/cce-operator/pkg/controller"
	"github.com/cnrancher/cce-operator/pkg/huawei/cce"
	"github.com/cnrancher/cce-operator/pkg/huawei/common"
	"github.com/cnrancher/cce-operator/pkg/huawei/eip"
	"github.com/cnrancher/cce-operator/pkg/utils"
	huawei_cce_model "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
	huawei_ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	huawei_ecs_model "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	huawei_ecs_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	huawei_eipv3 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3"
	huawei_eipv3_model "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3/model"
	huawei_eipv3_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3/region"
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/cloudcredentials"
	"github.com/rancher/shepherd/pkg/config"
	"github.com/sirupsen/logrus"
)

// CreateCCEHostedCluster is a helper function that creates an CCE hosted cluster
func CreateCCEHostedCluster(client *rancher.Client, displayName, cloudCredentialID string, enableClusterAlerting, enableClusterMonitoring, enableNetworkPolicy, windowsPreferedCluster bool, labels map[string]string) (*management.Cluster, error) {
	cceHostCluster := HostClusterConfig(displayName, cloudCredentialID)
	cluster := &management.Cluster{
		DockerRootDir:           "/var/lib/docker",
		CCEConfig:               cceHostCluster,
		Name:                    displayName,
		EnableClusterAlerting:   enableClusterAlerting,
		EnableClusterMonitoring: enableClusterMonitoring,
		EnableNetworkPolicy:     &enableNetworkPolicy,
		Labels:                  labels,
		WindowsPreferedCluster:  windowsPreferedCluster,
	}

	clusterResp, err := client.Management.Cluster.Create(cluster)
	if err != nil {
		return nil, err
	}
	return clusterResp, err
}

// UpdateNodePublicIP ensures the node has one public IP address.
func UpdateNodePublicIP(client *rancher.Client, ID string, cloudCredential *cloudcredentials.CloudCredential) (bool, error) {
	cluster, err := client.Management.Cluster.ByID(ID)
	if err != nil {
		return false, err
	}
	if cluster.CCEConfig == nil || len(cluster.CCEConfig.NodePools) == 0 {
		return false, nil
	}
	clusterID := cluster.CCEConfig.ClusterID
	if clusterID == "" {
		return false, nil
	}
	var c cloudcredentials.HuaweiCredentialConfig
	config.LoadConfig(cloudcredentials.HuaweiCredentialConfigurationFileKey, &c)
	auth := common.NewClientAuth(c.AccessKey, c.SecretKey, c.RegionID, c.ProjectID)
	driver := controller.NewHuaweiDriver(auth)
	nodesRes, err := cce.ListNodes(driver.CCE, clusterID)
	if err != nil {
		return false, err
	}
	if nodesRes == nil || nodesRes.Items == nil || len(*nodesRes.Items) == 0 {
		// waiting for cluster setup
		return false, nil
	}
	var count = 0
	for _, node := range *nodesRes.Items {
		if node.Status == nil || node.Status.Phase == nil || node.Status.ServerId == nil ||
			node.Metadata == nil || node.Metadata.Name == nil {
			continue
		}
		if *node.Status.Phase != huawei_cce_model.GetNodeStatusPhaseEnum().ACTIVE {
			logrus.Infof("waiting for node [%s] status %v",
				utils.Value(node.Metadata.Name), node.Status.Phase.Value())
			return false, nil
		}

		node, err := cce.ShowNode(driver.CCE, clusterID, utils.Value(node.Metadata.Uid))
		if err != nil {
			logrus.Errorf("ShowNode failed: %v", err)
			return false, err
		}
		if node == nil || node.Spec == nil || node.Status == nil || node.Metadata == nil {
			logrus.Errorf("ShowNode returns inalid data")
			return false, nil
		}
		if node.Status.PublicIP != nil && utils.Value(node.Status.PublicIP) != "" {
			logrus.Infof("node [%s] already have public IP [%s], skip",
				utils.Value(node.Metadata.Name), utils.Value(node.Status.PublicIP))
			count++
			continue
		}
		serverID := utils.Value(node.Status.ServerId)
		logrus.Infof("found node [%s] server ID [%s]", utils.Value(node.Metadata.Name), serverID)
		ecs := huawei_ecs.NewEcsClient(huawei_ecs.EcsClientBuilder().
			WithRegion(huawei_ecs_region.ValueOf(auth.Region)).
			WithCredential(auth.Credential).
			Build())
		listServerInterfacesRes, err := ecs.ListServerInterfaces(&huawei_ecs_model.ListServerInterfacesRequest{
			ServerId: serverID,
		})
		if err != nil {
			return false, err
		}
		if listServerInterfacesRes.InterfaceAttachments == nil || len(*listServerInterfacesRes.InterfaceAttachments) == 0 {
			logrus.Warnf("ecs server [%s] net interface is empty", serverID)
			continue
		}
		var portID string
		for _, i := range *listServerInterfacesRes.InterfaceAttachments {
			if i.PortId == nil {
				logrus.Warnf("InterfaceAttachments.PortId is nil pointer")
				continue
			}
			portID = utils.Value(i.PortId)
			break
		}
		if portID == "" {
			logrus.Warnf("failed to get ecs server [%s] net interface port id", serverID)
			continue
		}

		createPublicIPRes, err := eip.CreatePublicIP(driver.EIP, &ccev1.CCEEip{
			Iptype: cluster.CCEConfig.PublicIP.Eip.Iptype,
			Bandwidth: ccev1.CCEEipBandwidth{
				ChargeMode: cluster.CCEConfig.PublicIP.Eip.Bandwidth.ChargeMode,
				Size:       int32(cluster.CCEConfig.PublicIP.Eip.Bandwidth.Size),
				ShareType:  cluster.CCEConfig.PublicIP.Eip.Bandwidth.ShareType,
			},
		})
		if err != nil {
			return false, err
		}
		if createPublicIPRes == nil || createPublicIPRes.Publicip == nil ||
			createPublicIPRes.Publicip.Id == nil || createPublicIPRes.Publicip.PublicIpAddress == nil {
			logrus.Warnf("CreatePublicIP returns invalid data")
			continue
		}
		eipID := utils.Value(createPublicIPRes.Publicip.Id)
		eipAddr := utils.Value(createPublicIPRes.Publicip.PublicIpAddress)
		logrus.Infof("created EIP ID [%v]", eipID)

		eipv3 := huawei_eipv3.NewEipClient(
			huawei_eipv3.EipClientBuilder().
				WithRegion(huawei_eipv3_region.ValueOf(auth.Region)).
				WithCredential(auth.Credential).
				Build())
		_, err = eipv3.AssociatePublicips(&huawei_eipv3_model.AssociatePublicipsRequest{
			PublicipId: eipID,
			Body: &huawei_eipv3_model.AssociatePublicipsRequestBody{
				Publicip: &huawei_eipv3_model.AssociatePublicipsOption{
					AssociateInstanceType: huawei_eipv3_model.GetAssociatePublicipsOptionAssociateInstanceTypeEnum().PORT,
					AssociateInstanceId:   portID,
				},
			},
		})
		if err != nil {
			return false, err
		}
		logrus.Infof("successfully bind EIP [%v] for node [%v]",
			eipAddr, utils.Value(node.Metadata.Name))
		time.Sleep(5 * time.Second)
		count++
	}
	if count == len(*nodesRes.Items) {
		return true, nil
	}

	return false, nil
}
