package client

const (
	ClusterAdvancedSettingsType                         = "clusterAdvancedSettings"
	ClusterAdvancedSettingsFieldAsEnabled               = "asEnabled"
	ClusterAdvancedSettingsFieldAuditEnabled            = "auditEnabled"
	ClusterAdvancedSettingsFieldAuditLogTopicID         = "auditLogTopicId"
	ClusterAdvancedSettingsFieldAuditLogsetID           = "auditLogsetId"
	ClusterAdvancedSettingsFieldBasePodNumber           = "basePodNumber"
	ClusterAdvancedSettingsFieldCiliumMode              = "ciliumMode"
	ClusterAdvancedSettingsFieldContainerRuntime        = "containerRuntime"
	ClusterAdvancedSettingsFieldDeletionProtection      = "deletionProtection"
	ClusterAdvancedSettingsFieldEnableCustomizedPodCIDR = "enableCustomizedPodCIDR"
	ClusterAdvancedSettingsFieldEtcd                    = "etcd"
	ClusterAdvancedSettingsFieldIPVS                    = "ipvs"
	ClusterAdvancedSettingsFieldIsDualStack             = "isDualStack"
	ClusterAdvancedSettingsFieldIsNonStaticIpMode       = "isNonStaticIpMode"
	ClusterAdvancedSettingsFieldKubeAPIServer           = "kubeAPIServer"
	ClusterAdvancedSettingsFieldKubeControllerManager   = "kubeControllerManager"
	ClusterAdvancedSettingsFieldKubeProxyMode           = "kubeProxyMode"
	ClusterAdvancedSettingsFieldKubeScheduler           = "kubeScheduler"
	ClusterAdvancedSettingsFieldNetworkType             = "networkType"
	ClusterAdvancedSettingsFieldNodeNameType            = "nodeNameType"
	ClusterAdvancedSettingsFieldQGPUShareEnable         = "qgpuShareEnable"
	ClusterAdvancedSettingsFieldRuntimeVersion          = "runtimeVersion"
	ClusterAdvancedSettingsFieldVpcCniType              = "vpcCniType"
)

type ClusterAdvancedSettings struct {
	AsEnabled               bool     `json:"asEnabled,omitempty" yaml:"asEnabled,omitempty"`
	AuditEnabled            bool     `json:"auditEnabled,omitempty" yaml:"auditEnabled,omitempty"`
	AuditLogTopicID         string   `json:"auditLogTopicId,omitempty" yaml:"auditLogTopicId,omitempty"`
	AuditLogsetID           string   `json:"auditLogsetId,omitempty" yaml:"auditLogsetId,omitempty"`
	BasePodNumber           int64    `json:"basePodNumber,omitempty" yaml:"basePodNumber,omitempty"`
	CiliumMode              string   `json:"ciliumMode,omitempty" yaml:"ciliumMode,omitempty"`
	ContainerRuntime        string   `json:"containerRuntime,omitempty" yaml:"containerRuntime,omitempty"`
	DeletionProtection      bool     `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	EnableCustomizedPodCIDR bool     `json:"enableCustomizedPodCIDR,omitempty" yaml:"enableCustomizedPodCIDR,omitempty"`
	Etcd                    []string `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	IPVS                    bool     `json:"ipvs,omitempty" yaml:"ipvs,omitempty"`
	IsDualStack             bool     `json:"isDualStack,omitempty" yaml:"isDualStack,omitempty"`
	IsNonStaticIpMode       bool     `json:"isNonStaticIpMode,omitempty" yaml:"isNonStaticIpMode,omitempty"`
	KubeAPIServer           []string `json:"kubeAPIServer,omitempty" yaml:"kubeAPIServer,omitempty"`
	KubeControllerManager   []string `json:"kubeControllerManager,omitempty" yaml:"kubeControllerManager,omitempty"`
	KubeProxyMode           string   `json:"kubeProxyMode,omitempty" yaml:"kubeProxyMode,omitempty"`
	KubeScheduler           []string `json:"kubeScheduler,omitempty" yaml:"kubeScheduler,omitempty"`
	NetworkType             string   `json:"networkType,omitempty" yaml:"networkType,omitempty"`
	NodeNameType            string   `json:"nodeNameType,omitempty" yaml:"nodeNameType,omitempty"`
	QGPUShareEnable         bool     `json:"qgpuShareEnable,omitempty" yaml:"qgpuShareEnable,omitempty"`
	RuntimeVersion          string   `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`
	VpcCniType              string   `json:"vpcCniType,omitempty" yaml:"vpcCniType,omitempty"`
}
