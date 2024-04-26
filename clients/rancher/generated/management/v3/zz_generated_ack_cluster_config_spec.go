package client

const (
	ACKClusterConfigSpecType                          = "ackClusterConfigSpec"
	ACKClusterConfigSpecFieldAliyunCredentialSecret   = "aliyun_credential_secret"
	ACKClusterConfigSpecFieldCloudMonitorFlags        = "cloudMonitorFlags"
	ACKClusterConfigSpecFieldClusterID                = "cluster_id"
	ACKClusterConfigSpecFieldClusterIsUpgrading       = "clusterIsUpgrading"
	ACKClusterConfigSpecFieldClusterType              = "clusterType"
	ACKClusterConfigSpecFieldContainerCidr            = "containerCidr"
	ACKClusterConfigSpecFieldDisableRollback          = "disableRollback"
	ACKClusterConfigSpecFieldEndpointPublicAccess     = "endpointPublicAccess"
	ACKClusterConfigSpecFieldImported                 = "imported"
	ACKClusterConfigSpecFieldKeyPair                  = "keyPair"
	ACKClusterConfigSpecFieldKubernetesVersion        = "kubernetesVersion"
	ACKClusterConfigSpecFieldLoginPassword            = "loginPassword"
	ACKClusterConfigSpecFieldMasterAutoRenew          = "masterAutoRenew"
	ACKClusterConfigSpecFieldMasterAutoRenewPeriod    = "masterAutoRenewPeriod"
	ACKClusterConfigSpecFieldMasterCount              = "masterCount"
	ACKClusterConfigSpecFieldMasterInstanceChargeType = "masterInstanceChargeType"
	ACKClusterConfigSpecFieldMasterInstanceTypes      = "masterInstanceTypes"
	ACKClusterConfigSpecFieldMasterPeriod             = "masterPeriod"
	ACKClusterConfigSpecFieldMasterPeriodUnit         = "masterPeriodUnit"
	ACKClusterConfigSpecFieldMasterSystemDiskCategory = "masterSystemDiskCategory"
	ACKClusterConfigSpecFieldMasterSystemDiskSize     = "masterSystemDiskSize"
	ACKClusterConfigSpecFieldMasterVswitchIds         = "masterVswitchIds"
	ACKClusterConfigSpecFieldName                     = "name"
	ACKClusterConfigSpecFieldNodeCidrMask             = "nodeCidrMask"
	ACKClusterConfigSpecFieldNodePoolList             = "node_pool_list"
	ACKClusterConfigSpecFieldOsType                   = "osType"
	ACKClusterConfigSpecFieldPauseClusterUpgrade      = "pauseClusterUpgrade"
	ACKClusterConfigSpecFieldPlatform                 = "platform"
	ACKClusterConfigSpecFieldProxyMode                = "proxyMode"
	ACKClusterConfigSpecFieldRegionID                 = "regionId"
	ACKClusterConfigSpecFieldResourceGroupID          = "resourceGroupId"
	ACKClusterConfigSpecFieldSSHFlags                 = "sshFlags"
	ACKClusterConfigSpecFieldSecurityGroupID          = "securityGroupId"
	ACKClusterConfigSpecFieldServiceCidr              = "serviceCidr"
	ACKClusterConfigSpecFieldSnatEntry                = "snatEntry"
	ACKClusterConfigSpecFieldTimeoutMins              = "timeoutMins"
	ACKClusterConfigSpecFieldVpcID                    = "vpcId"
	ACKClusterConfigSpecFieldVswitchIds               = "vswitchIds"
)

type ACKClusterConfigSpec struct {
	AliyunCredentialSecret   string         `json:"aliyun_credential_secret,omitempty" yaml:"aliyun_credential_secret,omitempty"`
	CloudMonitorFlags        bool           `json:"cloudMonitorFlags,omitempty" yaml:"cloudMonitorFlags,omitempty"`
	ClusterID                string         `json:"cluster_id,omitempty" yaml:"cluster_id,omitempty"`
	ClusterIsUpgrading       bool           `json:"clusterIsUpgrading,omitempty" yaml:"clusterIsUpgrading,omitempty"`
	ClusterType              string         `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`
	ContainerCidr            string         `json:"containerCidr,omitempty" yaml:"containerCidr,omitempty"`
	DisableRollback          bool           `json:"disableRollback,omitempty" yaml:"disableRollback,omitempty"`
	EndpointPublicAccess     bool           `json:"endpointPublicAccess,omitempty" yaml:"endpointPublicAccess,omitempty"`
	Imported                 bool           `json:"imported,omitempty" yaml:"imported,omitempty"`
	KeyPair                  string         `json:"keyPair,omitempty" yaml:"keyPair,omitempty"`
	KubernetesVersion        string         `json:"kubernetesVersion,omitempty" yaml:"kubernetesVersion,omitempty"`
	LoginPassword            string         `json:"loginPassword,omitempty" yaml:"loginPassword,omitempty"`
	MasterAutoRenew          bool           `json:"masterAutoRenew,omitempty" yaml:"masterAutoRenew,omitempty"`
	MasterAutoRenewPeriod    int64          `json:"masterAutoRenewPeriod,omitempty" yaml:"masterAutoRenewPeriod,omitempty"`
	MasterCount              int64          `json:"masterCount,omitempty" yaml:"masterCount,omitempty"`
	MasterInstanceChargeType string         `json:"masterInstanceChargeType,omitempty" yaml:"masterInstanceChargeType,omitempty"`
	MasterInstanceTypes      []string       `json:"masterInstanceTypes,omitempty" yaml:"masterInstanceTypes,omitempty"`
	MasterPeriod             int64          `json:"masterPeriod,omitempty" yaml:"masterPeriod,omitempty"`
	MasterPeriodUnit         string         `json:"masterPeriodUnit,omitempty" yaml:"masterPeriodUnit,omitempty"`
	MasterSystemDiskCategory string         `json:"masterSystemDiskCategory,omitempty" yaml:"masterSystemDiskCategory,omitempty"`
	MasterSystemDiskSize     int64          `json:"masterSystemDiskSize,omitempty" yaml:"masterSystemDiskSize,omitempty"`
	MasterVswitchIds         []string       `json:"masterVswitchIds,omitempty" yaml:"masterVswitchIds,omitempty"`
	Name                     string         `json:"name,omitempty" yaml:"name,omitempty"`
	NodeCidrMask             int64          `json:"nodeCidrMask,omitempty" yaml:"nodeCidrMask,omitempty"`
	NodePoolList             []NodePoolInfo `json:"node_pool_list,omitempty" yaml:"node_pool_list,omitempty"`
	OsType                   string         `json:"osType,omitempty" yaml:"osType,omitempty"`
	PauseClusterUpgrade      bool           `json:"pauseClusterUpgrade,omitempty" yaml:"pauseClusterUpgrade,omitempty"`
	Platform                 string         `json:"platform,omitempty" yaml:"platform,omitempty"`
	ProxyMode                string         `json:"proxyMode,omitempty" yaml:"proxyMode,omitempty"`
	RegionID                 string         `json:"regionId,omitempty" yaml:"regionId,omitempty"`
	ResourceGroupID          string         `json:"resourceGroupId,omitempty" yaml:"resourceGroupId,omitempty"`
	SSHFlags                 bool           `json:"sshFlags,omitempty" yaml:"sshFlags,omitempty"`
	SecurityGroupID          string         `json:"securityGroupId,omitempty" yaml:"securityGroupId,omitempty"`
	ServiceCidr              string         `json:"serviceCidr,omitempty" yaml:"serviceCidr,omitempty"`
	SnatEntry                bool           `json:"snatEntry,omitempty" yaml:"snatEntry,omitempty"`
	TimeoutMins              int64          `json:"timeoutMins,omitempty" yaml:"timeoutMins,omitempty"`
	VpcID                    string         `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
	VswitchIds               []string       `json:"vswitchIds,omitempty" yaml:"vswitchIds,omitempty"`
}
