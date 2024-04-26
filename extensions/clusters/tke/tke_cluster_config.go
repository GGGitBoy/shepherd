package tke

import (
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/pkg/config"
)

const (
	// The json/yaml config key for the TKE hosted cluster config
	TKEClusterConfigConfigurationFileKey = "tkeClusterConfig"
)

// ClusterConfig is the configuration needed to create an TKE host cluster
type ClusterConfig struct {
	Region                  string                   `json:"region,omitempty"`
	ClusterID               string                   `json:"clusterId,omitempty"`
	ClusterEndpoint         *ClusterEndpoint         `json:"clusterEndpoint,omitempty"`
	ClusterBasicSettings    *ClusterBasicSettings    `json:"clusterBasicSettings,omitempty"`
	ClusterCIDRSettings     *ClusterCIDRSettings     `json:"clusterCIDRSettings,omitempty"`
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"clusterAdvancedSettings,omitempty"`
	ExtensionAddon          []ExtensionAddon         `json:"extensionAddon,omitempty"`
	RunInstancesForNode     *RunInstancesForNode     `json:"runInstancesForNode,omitempty"`
	NodePoolList            []NodePoolDetail         `json:"nodePoolList,omitempty"`
}

type ExtensionAddon struct {
	AddonName  string `json:"addonName,omitempty"`
	AddonParam string `json:"addonParam,omitempty"`
}

type ClusterEndpoint struct {
	Enable              bool   `json:"enable,omitempty"`
	Domain              string `json:"domain,omitempty"`
	SubnetID            string `json:"subnetId,omitempty"`
	ExtensiveParameters string `json:"extensiveParameters,omitempty"`
	SecurityGroup       string `json:"securityGroup,omitempty"`
}

type NodePoolDetail struct {
	ClusterID            string               `json:"clusterId,omitempty"`
	NodePoolID           string               `json:"nodePoolId,omitempty"`
	AutoScalingGroupPara AutoScalingGroupPara `json:"autoScalingGroupPara,omitempty"`
	LaunchConfigurePara  LaunchConfigurePara  `json:"launchConfigurePara,omitempty"`
	EnableAutoscale      bool                 `json:"enableAutoscale,omitempty"`
	Name                 string               `json:"name,omitempty"`
	Labels               []string             `json:"labels,omitempty"`
	Taints               []string             `json:"taints,omitempty"`
	NodePoolOs           string               `json:"nodePoolOs,omitempty"`
	OsCustomizeType      string               `json:"osCustomizeType,omitempty"`
	Tags                 []string             `json:"tags,omitempty"`
	DeletionProtection   bool                 `json:"deletionProtection,omitempty"`
}

type AutoScalingGroupPara struct {
	AutoScalingGroupName string   `json:"autoScalingGroupName,omitempty"`
	MaxSize              int64    `json:"maxSize,omitempty"`
	MinSize              int64    `json:"minSize,omitempty"`
	DesiredCapacity      int64    `json:"desiredCapacity,omitempty"`
	VpcID                string   `json:"vpcId,omitempty"`
	SubnetIDs            []string `json:"subnetIds,omitempty"`
}

type LaunchConfigurePara struct {
	LaunchConfigurationName string     `json:"launchConfigurationName,omitempty"`
	InstanceType            string     `json:"instanceType,omitempty"`
	SystemDisk              DataDisk   `json:"systemDisk,omitempty"`
	InternetChargeType      string     `json:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int64      `json:"internetMaxBandwidthOut,omitempty"`
	PublicIPAssigned        bool       `json:"publicIpAssigned,omitempty"`
	DataDisks               []DataDisk `json:"dataDisks,omitempty"`
	KeyIDs                  []string   `json:"keyIds,omitempty"`
	SecurityGroupIDs        []string   `json:"securityGroupIds,omitempty"`
	InstanceChargeType      string     `json:"instanceChargeType,omitempty"`
}

type ClusterBasicSettings struct {
	ClusterType        string   `json:"clusterType,omitempty"`
	ClusterOs          string   `json:"clusterOs,omitempty"`
	ClusterVersion     string   `json:"clusterVersion,omitempty"`
	ClusterName        string   `json:"clusterName,omitempty"`
	ClusterDescription string   `json:"clusterDescription,omitempty"`
	VpcID              string   `json:"vpcId,omitempty"`
	ProjectID          int64    `json:"projectId,omitempty"`
	Tags               []string `json:"tags,omitempty"`
	ClusterLevel       string   `json:"clusterLevel,omitempty"`
	IsAutoUpgrade      bool     `json:"isAutoUpgrade,omitempty"`
}

type ClusterCIDRSettings struct {
	ClusterCIDR               string   `json:"clusterCIDR,omitempty"`
	IgnoreClusterCIDRConflict bool     `json:"ignoreClusterCIDRConflict,omitempty"`
	MaxNodePodNum             int64    `json:"maxNodePodNum,omitempty"`
	MaxClusterServiceNum      int64    `json:"maxClusterServiceNum,omitempty"`
	ServiceCIDR               string   `json:"serviceCIDR,omitempty"`
	EniSubnetIDs              []string `json:"eniSubnetIds,omitempty"`
	ClaimExpiredSeconds       int64    `json:"claimExpiredSeconds,omitempty"`
	IgnoreServiceCIDRConflict bool     `json:"ignoreServiceCIDRConflict,omitempty"`
	OsCustomizeType           string   `json:"osCustomizeType,omitempty"`
	SubnetID                  string   `json:"subnetId,omitempty"`
}

type ClusterAdvancedSettings struct {
	IPVS                    bool     `json:"ipvs,omitempty"`
	AsEnabled               bool     `json:"asEnabled,omitempty"`
	ContainerRuntime        string   `json:"containerRuntime,omitempty"`
	NodeNameType            string   `json:"nodeNameType,omitempty"`
	KubeAPIServer           []string `json:"kubeAPIServer,omitempty"`
	KubeControllerManager   []string `json:"kubeControllerManager,omitempty"`
	KubeScheduler           []string `json:"kubeScheduler,omitempty"`
	Etcd                    []string `json:"etcd,omitempty"`
	NetworkType             string   `json:"networkType,omitempty"`
	IsNonStaticIPMode       bool     `json:"isNonStaticIpMode,omitempty"`
	DeletionProtection      bool     `json:"deletionProtection,omitempty"`
	KubeProxyMode           string   `json:"kubeProxyMode,omitempty"`
	AuditEnabled            bool     `json:"auditEnabled,omitempty"`
	AuditLogsetID           string   `json:"auditLogsetId,omitempty"`
	AuditLogTopicID         string   `json:"auditLogTopicId,omitempty"`
	VpcCniType              string   `json:"vpcCniType,omitempty"`
	RuntimeVersion          string   `json:"runtimeVersion,omitempty"`
	EnableCustomizedPodCIDR bool     `json:"enableCustomizedPodCIDR,omitempty"`
	BasePodNumber           int64    `json:"basePodNumber,omitempty"`
	CiliumMode              string   `json:"ciliumMode,omitempty"`
	IsDualStack             bool     `json:"isDualStack,omitempty"`
	QGPUShareEnable         bool     `json:"qgpuShareEnable,omitempty"`
}

type RunInstancesForNode struct {
	NodeRole                string   `json:"nodeRole,omitempty"`
	InstanceChargeType      string   `json:"instanceChargeType,omitempty"`
	Zone                    string   `json:"zone,omitempty"`
	InstanceCount           int64    `json:"instanceCount,omitempty"`
	ProjectID               int64    `json:"projectId,omitempty"`
	InstanceType            string   `json:"instanceType,omitempty"`
	ImageID                 string   `json:"imageId,omitempty"`
	SystemDisk              DataDisk `json:"systemDisk,omitempty"`
	VpcID                   string   `json:"vpcId,omitempty"`
	SubnetID                string   `json:"subnetId,omitempty"`
	InternetChargeType      string   `json:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int64    `json:"internetMaxBandwidthOut,omitempty"`
	PublicIPAssigned        bool     `json:"publicIpAssigned,omitempty"`
	InstanceName            string   `json:"instanceName,omitempty"`
	KeyIDs                  []string `json:"keyIds,omitempty"`
	SecurityService         bool     `json:"securityService,omitempty"`
	MonitorService          bool     `json:"monitorService,omitempty"`
	UserData                string   `json:"userData,omitempty"`
}

type DataDisk struct {
	DiskSize int64  `json:"diskSize,omitempty"`
	DiskType string `json:"diskType,omitempty"`
}

func nodeDataVolumesConstructor(nodeVolumesConfig []DataDisk) []management.DataDisk {
	var dataVolumes []management.DataDisk
	for _, dv := range nodeVolumesConfig {
		dataVolumes = append(dataVolumes, management.DataDisk{
			DiskSize: dv.DiskSize,
			DiskType: dv.DiskType,
		})
	}
	return dataVolumes
}

func nodeDataVolumeConstructor(nodeVolumesConfig DataDisk) *management.DataDisk {
	return &management.DataDisk{
		DiskSize: nodeVolumesConfig.DiskSize,
		DiskType: nodeVolumesConfig.DiskType,
	}
}

func nodePoolsConstructor(nodePoolsConfig []NodePoolDetail) []management.NodePoolDetail {
	var nodePoolList []management.NodePoolDetail
	for _, nodePool := range nodePoolsConfig {
		nodePoolList = append(nodePoolList, management.NodePoolDetail{
			ClusterID:  nodePool.ClusterID,
			NodePoolID: nodePool.NodePoolID,
			AutoScalingGroupPara: &management.AutoScalingGroupPara{
				AutoScalingGroupName: nodePool.AutoScalingGroupPara.AutoScalingGroupName,
				MaxSize:              nodePool.AutoScalingGroupPara.MaxSize,
				MinSize:              nodePool.AutoScalingGroupPara.MinSize,
				DesiredCapacity:      nodePool.AutoScalingGroupPara.DesiredCapacity,
				VpcID:                nodePool.AutoScalingGroupPara.VpcID,
				SubnetIDs:            nodePool.AutoScalingGroupPara.SubnetIDs,
			},

			LaunchConfigurePara: &management.LaunchConfigurePara{
				LaunchConfigurationName: nodePool.LaunchConfigurePara.LaunchConfigurationName,
				InstanceType:            nodePool.LaunchConfigurePara.InstanceType,
				SystemDisk:              nodeDataVolumeConstructor(nodePool.LaunchConfigurePara.SystemDisk),
				InternetChargeType:      nodePool.LaunchConfigurePara.InternetChargeType,
				InternetMaxBandwidthOut: nodePool.LaunchConfigurePara.InternetMaxBandwidthOut,
				PublicIpAssigned:        nodePool.LaunchConfigurePara.PublicIPAssigned,
				DataDisks:               nodeDataVolumesConstructor(nodePool.LaunchConfigurePara.DataDisks),
				KeyIDs:                  nodePool.LaunchConfigurePara.KeyIDs,
				SecurityGroupIDs:        nodePool.LaunchConfigurePara.SecurityGroupIDs,
				InstanceChargeType:      nodePool.LaunchConfigurePara.InstanceChargeType,
			},
			Name:               nodePool.Name,
			Labels:             nodePool.Labels,
			Taints:             nodePool.Taints,
			NodePoolOs:         nodePool.NodePoolOs,
			OsCustomizeType:    nodePool.OsCustomizeType,
			Tags:               nodePool.Tags,
			DeletionProtection: nodePool.DeletionProtection,
		})
	}

	return nodePoolList
}

func HostClusterConfig(name, cloudCredentialID string) *management.TKEClusterConfigSpec {
	var tkeClusterConfig ClusterConfig
	config.LoadConfig(TKEClusterConfigConfigurationFileKey, &tkeClusterConfig)

	return &management.TKEClusterConfigSpec{
		Region:              tkeClusterConfig.Region,
		TKECredentialSecret: cloudCredentialID,
		ClusterEndpoint: &management.ClusterEndpoint{
			Domain:              tkeClusterConfig.ClusterEndpoint.Domain,
			Enable:              tkeClusterConfig.ClusterEndpoint.Enable,
			ExtensiveParameters: tkeClusterConfig.ClusterEndpoint.ExtensiveParameters,
			SecurityGroup:       tkeClusterConfig.ClusterEndpoint.SecurityGroup,
			SubnetID:            tkeClusterConfig.ClusterEndpoint.SubnetID,
		},

		ClusterBasicSettings: &management.ClusterBasicSettings{
			ClusterType:        tkeClusterConfig.ClusterBasicSettings.ClusterType,
			ClusterOs:          tkeClusterConfig.ClusterBasicSettings.ClusterOs,
			ClusterVersion:     tkeClusterConfig.ClusterBasicSettings.ClusterVersion,
			ClusterName:        tkeClusterConfig.ClusterBasicSettings.ClusterName,
			ClusterDescription: tkeClusterConfig.ClusterBasicSettings.ClusterDescription,
			VpcID:              tkeClusterConfig.ClusterBasicSettings.VpcID,
			Tags:               tkeClusterConfig.ClusterBasicSettings.Tags,
			ClusterLevel:       tkeClusterConfig.ClusterBasicSettings.ClusterLevel,
			IsAutoUpgrade:      tkeClusterConfig.ClusterBasicSettings.IsAutoUpgrade,
			ProjectID:          tkeClusterConfig.ClusterBasicSettings.ProjectID,
		},

		ClusterCIDRSettings: &management.ClusterCIDRSettings{
			ClusterCIDR:               tkeClusterConfig.ClusterCIDRSettings.ClusterCIDR,
			IgnoreClusterCIDRConflict: tkeClusterConfig.ClusterCIDRSettings.IgnoreClusterCIDRConflict,
			MaxNodePodNum:             tkeClusterConfig.ClusterCIDRSettings.MaxNodePodNum,
			MaxClusterServiceNum:      tkeClusterConfig.ClusterCIDRSettings.MaxClusterServiceNum,
			ServiceCIDR:               tkeClusterConfig.ClusterCIDRSettings.ServiceCIDR,
			EniSubnetIDs:              tkeClusterConfig.ClusterCIDRSettings.EniSubnetIDs,
			IgnoreServiceCIDRConflict: tkeClusterConfig.ClusterCIDRSettings.IgnoreServiceCIDRConflict,
			OsCustomizeType:           tkeClusterConfig.ClusterCIDRSettings.OsCustomizeType,
		},

		ClusterAdvancedSettings: &management.ClusterAdvancedSettings{
			IPVS:             tkeClusterConfig.ClusterAdvancedSettings.IPVS,
			ContainerRuntime: tkeClusterConfig.ClusterAdvancedSettings.ContainerRuntime,
			RuntimeVersion:   tkeClusterConfig.ClusterAdvancedSettings.RuntimeVersion,
			QGPUShareEnable:  tkeClusterConfig.ClusterAdvancedSettings.QGPUShareEnable,
		},

		NodePoolList: nodePoolsConstructor(tkeClusterConfig.NodePoolList),
	}
}
