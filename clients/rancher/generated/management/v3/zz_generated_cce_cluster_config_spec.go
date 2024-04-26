package client

const (
	CCEClusterConfigSpecType                        = "cceClusterConfigSpec"
	CCEClusterConfigSpecFieldAuthentication         = "authentication"
	CCEClusterConfigSpecFieldBillingMode            = "clusterBillingMode"
	CCEClusterConfigSpecFieldCategory               = "category"
	CCEClusterConfigSpecFieldClusterID              = "clusterID"
	CCEClusterConfigSpecFieldContainerNetwork       = "containerNetwork"
	CCEClusterConfigSpecFieldCreatedNodePoolIDs     = "createdNodePoolIDs"
	CCEClusterConfigSpecFieldDescription            = "description"
	CCEClusterConfigSpecFieldEniNetwork             = "eniNetwork"
	CCEClusterConfigSpecFieldExtendParam            = "extendParam"
	CCEClusterConfigSpecFieldFlavor                 = "flavor"
	CCEClusterConfigSpecFieldHostNetwork            = "hostNetwork"
	CCEClusterConfigSpecFieldHuaweiCredentialSecret = "huaweiCredentialSecret"
	CCEClusterConfigSpecFieldImported               = "imported"
	CCEClusterConfigSpecFieldIpv6Enable             = "ipv6Enable"
	CCEClusterConfigSpecFieldKubeProxyMode          = "kubeProxyMode"
	CCEClusterConfigSpecFieldKubernetesSvcIPRange   = "kubernetesSvcIPRange"
	CCEClusterConfigSpecFieldLabels                 = "labels"
	CCEClusterConfigSpecFieldName                   = "name"
	CCEClusterConfigSpecFieldNatGateway             = "natGateway"
	CCEClusterConfigSpecFieldNodePools              = "nodePools"
	CCEClusterConfigSpecFieldPublicAccess           = "publicAccess"
	CCEClusterConfigSpecFieldPublicIP               = "publicIP"
	CCEClusterConfigSpecFieldRegionID               = "regionID"
	CCEClusterConfigSpecFieldTags                   = "tags"
	CCEClusterConfigSpecFieldType                   = "type"
	CCEClusterConfigSpecFieldVersion                = "version"
)

type CCEClusterConfigSpec struct {
	Authentication         *CCEAuthentication     `json:"authentication,omitempty" yaml:"authentication,omitempty"`
	BillingMode            int64                  `json:"clusterBillingMode,omitempty" yaml:"clusterBillingMode,omitempty"`
	Category               string                 `json:"category,omitempty" yaml:"category,omitempty"`
	ClusterID              string                 `json:"clusterID,omitempty" yaml:"clusterID,omitempty"`
	ContainerNetwork       *CCEContainerNetwork   `json:"containerNetwork,omitempty" yaml:"containerNetwork,omitempty"`
	CreatedNodePoolIDs     map[string]string      `json:"createdNodePoolIDs,omitempty" yaml:"createdNodePoolIDs,omitempty"`
	Description            string                 `json:"description,omitempty" yaml:"description,omitempty"`
	EniNetwork             *CCEEniNetwork         `json:"eniNetwork,omitempty" yaml:"eniNetwork,omitempty"`
	ExtendParam            *CCEClusterExtendParam `json:"extendParam,omitempty" yaml:"extendParam,omitempty"`
	Flavor                 string                 `json:"flavor,omitempty" yaml:"flavor,omitempty"`
	HostNetwork            *CCEHostNetwork        `json:"hostNetwork,omitempty" yaml:"hostNetwork,omitempty"`
	HuaweiCredentialSecret string                 `json:"huaweiCredentialSecret,omitempty" yaml:"huaweiCredentialSecret,omitempty"`
	Imported               bool                   `json:"imported,omitempty" yaml:"imported,omitempty"`
	Ipv6Enable             bool                   `json:"ipv6Enable,omitempty" yaml:"ipv6Enable,omitempty"`
	KubeProxyMode          string                 `json:"kubeProxyMode,omitempty" yaml:"kubeProxyMode,omitempty"`
	KubernetesSvcIPRange   string                 `json:"kubernetesSvcIPRange,omitempty" yaml:"kubernetesSvcIPRange,omitempty"`
	Labels                 map[string]string      `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                   string                 `json:"name,omitempty" yaml:"name,omitempty"`
	NatGateway             *CCENatGateway         `json:"natGateway,omitempty" yaml:"natGateway,omitempty"`
	NodePools              []CCENodePool          `json:"nodePools,omitempty" yaml:"nodePools,omitempty"`
	PublicAccess           bool                   `json:"publicAccess,omitempty" yaml:"publicAccess,omitempty"`
	PublicIP               *CCEClusterPublicIP    `json:"publicIP,omitempty" yaml:"publicIP,omitempty"`
	RegionID               string                 `json:"regionID,omitempty" yaml:"regionID,omitempty"`
	Tags                   map[string]string      `json:"tags,omitempty" yaml:"tags,omitempty"`
	Type                   string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Version                string                 `json:"version,omitempty" yaml:"version,omitempty"`
}
