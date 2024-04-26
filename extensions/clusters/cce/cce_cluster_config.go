package cce

import (
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/pkg/config"
)

const (
	// The json/yaml config key for the CCE hosted cluster config
	CCEClusterConfigConfigurationFileKey = "cceClusterConfig"
)

// ClusterConfig is the configuration needed to create an CCE host cluster
type ClusterConfig struct {
	RegionID             string            `json:"regionID" yaml:"regionID"`
	Name                 string            `json:"name" yaml:"name"`
	Labels               map[string]string `json:"labels" yaml:"labels"`
	Flavor               string            `json:"flavor" yaml:"flavor"`
	Version              string            `json:"version" yaml:"version"`
	Description          string            `json:"description" yaml:"description"`
	HostNetwork          HostNetwork       `json:"hostNetwork" yaml:"hostNetwork"`
	ContainerNetwork     ContainerNetwork  `json:"containerNetwork" yaml:"containerNetwork"`
	Authentication       Authentication    `json:"authentication" yaml:"authentication"`
	BillingMode          int64             `json:"clusterBillingMode" yaml:"clusterBillingMode"`
	KubernetesSvcIPRange string            `json:"kubernetesSvcIPRange" yaml:"kubernetesSvcIPRange"`
	Tags                 map[string]string `json:"tags" yaml:"tags"`
	KubeProxyMode        string            `json:"kubeProxyMode" yaml:"kubeProxyMode"`
	PublicAccess         bool              `json:"publicAccess" yaml:"publicAccess"`
	PublicIP             ClusterPublicIP   `json:"publicIP" yaml:"publicIP"`
	NodePools            []NodePool        `json:"nodePools" yaml:"nodePools"`
}

type HostNetwork struct {
	VpcID         string `json:"vpcID" yaml:"vpcID"`
	SubnetID      string `json:"subnetID" yaml:"subnetID"`
	SecurityGroup string `json:"securityGroup" yaml:"securityGroup"`
}

type ContainerNetwork struct {
	Mode string `json:"mode" yaml:"mode"`
	CIDR string `json:"cidr" yaml:"cidr"`
}

type Authentication struct {
	Mode                string              `json:"mode" yaml:"mode"`
	AuthenticatingProxy AuthenticatingProxy `json:"authenticatingProxy" yaml:"authenticatingProxy"`
}

type AuthenticatingProxy struct {
	Ca         string `json:"ca" yaml:"ca"`
	Cert       string `json:"cert" yaml:"cert"`
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
}

type NodePool struct {
	Name                 string       `json:"name" yaml:"name"`
	Type                 string       `json:"type" yaml:"type"`
	NodeTemplate         NodeTemplate `json:"nodeTemplate" yaml:"nodeTemplate"`
	InitialNodeCount     int64        `json:"initialNodeCount" yaml:"initialNodeCount"`
	PodSecurityGroups    []string     `json:"podSecurityGroups" yaml:"podSecurityGroups"`
	CustomSecurityGroups []string     `json:"customSecurityGroups" yaml:"customSecurityGroups"`
}

type NodeTemplate struct {
	Flavor          string       `json:"flavor" yaml:"flavor"`
	AvailableZone   string       `json:"availableZone" yaml:"availableZone"`
	OperatingSystem string       `json:"operatingSystem" yaml:"operatingSystem"`
	SSHKey          string       `json:"sshKey" yaml:"sshKey"`
	RootVolume      NodeVolume   `json:"rootVolume" yaml:"rootVolume"`
	DataVolumes     []NodeVolume `json:"dataVolumes" yaml:"dataVolumes"`
	BillingMode     int64        `json:"billingMode" yaml:"billingMode"`
	Runtime         string       `json:"runtime" yaml:"runtime"`
}

type NodeVolume struct {
	Size int64  `json:"size" yaml:"size"`
	Type string `json:"type" yaml:"type"`
}

type EipBandwidth struct {
	ChargeMode string `json:"chargeMode" yaml:"chargeMode"`
	Size       int64  `json:"size" yaml:"size"`
	ShareType  string `json:"shareType" yaml:"shareType"`
}

type Eip struct {
	Iptype    string       `json:"ipType" yaml:"ipType"`
	Bandwidth EipBandwidth `json:"bandwidth" yaml:"bandwidth"`
}

type ClusterPublicIP struct {
	CreateEIP bool `json:"createEIP" yaml:"createEIP"`
	Eip       Eip  `json:"eip" yaml:"eip"`
}

func nodeDataVolumesConstructor(nodeVolumesConfig []NodeVolume) []management.CCENodeVolume {
	var dataVolumes []management.CCENodeVolume
	for _, dv := range nodeVolumesConfig {
		dataVolumes = append(dataVolumes, management.CCENodeVolume{
			Size: dv.Size,
			Type: dv.Type,
		})
	}
	return dataVolumes
}

func nodePoolsConstructor(nodePoolsConfig []NodePool) []management.CCENodePool {
	var nodePools []management.CCENodePool
	for _, np := range nodePoolsConfig {
		nodePools = append(nodePools, management.CCENodePool{
			Autoscaling:          nil,
			CustomSecurityGroups: nil,
			ID:                   "",
			InitialNodeCount:     np.InitialNodeCount,
			Name:                 np.Name,
			Type:                 np.Type,
			NodeTemplate: &management.CCENodeTemplate{
				AvailableZone:   np.NodeTemplate.AvailableZone,
				BillingMode:     np.NodeTemplate.BillingMode,
				DataVolumes:     nodeDataVolumesConstructor(np.NodeTemplate.DataVolumes),
				ExtendParam:     nil,
				Flavor:          np.NodeTemplate.Flavor,
				OperatingSystem: np.NodeTemplate.OperatingSystem,
				PublicIP:        nil,
				RootVolume: &management.CCENodeVolume{
					Size: np.NodeTemplate.RootVolume.Size,
					Type: np.NodeTemplate.RootVolume.Type,
				},
				Runtime: np.NodeTemplate.Runtime,
				SSHKey:  np.NodeTemplate.SSHKey,
			},
		})
	}
	return nodePools
}

func HostClusterConfig(name, cloudCredentialID string) *management.CCEClusterConfigSpec {
	var cceClusterConfig ClusterConfig
	config.LoadConfig(CCEClusterConfigConfigurationFileKey, &cceClusterConfig)

	return &management.CCEClusterConfigSpec{
		Authentication: &management.CCEAuthentication{
			AuthenticatingProxy: &management.CCEAuthenticatingProxy{
				Ca:   cceClusterConfig.Authentication.AuthenticatingProxy.Ca,
				Cert: cceClusterConfig.Authentication.AuthenticatingProxy.Cert,
			},
			Mode: cceClusterConfig.Authentication.Mode,
		},
		BillingMode: 0,
		Category:    "CCE",
		ClusterID:   "",
		ContainerNetwork: &management.CCEContainerNetwork{
			CIDR: cceClusterConfig.ContainerNetwork.CIDR,
			Mode: cceClusterConfig.ContainerNetwork.Mode,
		},
		CreatedNodePoolIDs: nil,
		Description:        cceClusterConfig.Description,
		EniNetwork:         nil,
		ExtendParam:        nil,
		Flavor:             cceClusterConfig.Flavor,
		Version:            cceClusterConfig.Version,
		HostNetwork: &management.CCEHostNetwork{
			SecurityGroup: cceClusterConfig.HostNetwork.SecurityGroup,
			SubnetID:      cceClusterConfig.HostNetwork.SubnetID,
			VpcID:         cceClusterConfig.HostNetwork.VpcID,
		},
		HuaweiCredentialSecret: cloudCredentialID,
		Imported:               false,
		Ipv6Enable:             false,
		KubeProxyMode:          cceClusterConfig.KubeProxyMode,
		KubernetesSvcIPRange:   cceClusterConfig.KubernetesSvcIPRange,
		Labels:                 cceClusterConfig.Labels,
		Name:                   name,
		NodePools:              nodePoolsConstructor(cceClusterConfig.NodePools),
		PublicAccess:           cceClusterConfig.PublicAccess,
		PublicIP: &management.CCEClusterPublicIP{
			CreateEIP: cceClusterConfig.PublicIP.CreateEIP,
			Eip: &management.CCEEip{
				Bandwidth: &management.CCEEipBandwidth{
					ChargeMode: cceClusterConfig.PublicIP.Eip.Bandwidth.ChargeMode,
					ShareType:  cceClusterConfig.PublicIP.Eip.Bandwidth.ShareType,
					Size:       cceClusterConfig.PublicIP.Eip.Bandwidth.Size,
				},
				Iptype: cceClusterConfig.PublicIP.Eip.Iptype,
			},
		},
		RegionID: cceClusterConfig.RegionID,
		Tags:     cceClusterConfig.Tags,
		Type:     "CCE",
	}
}
