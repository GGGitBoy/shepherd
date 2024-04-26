package client

const (
	TKEClusterConfigSpecType                         = "tkeClusterConfigSpec"
	TKEClusterConfigSpecFieldClusterAdvancedSettings = "clusterAdvancedSettings"
	TKEClusterConfigSpecFieldClusterBasicSettings    = "clusterBasicSettings"
	TKEClusterConfigSpecFieldClusterCIDRSettings     = "clusterCIDRSettings"
	TKEClusterConfigSpecFieldClusterEndpoint         = "clusterEndpoint"
	TKEClusterConfigSpecFieldClusterID               = "clusterId"
	TKEClusterConfigSpecFieldExtensionAddon          = "extensionAddon"
	TKEClusterConfigSpecFieldImported                = "imported"
	TKEClusterConfigSpecFieldNodePoolList            = "nodePoolList"
	TKEClusterConfigSpecFieldRegion                  = "region"
	TKEClusterConfigSpecFieldRunInstancesForNode     = "runInstancesForNode"
	TKEClusterConfigSpecFieldTKECredentialSecret     = "tkeCredentialSecret"
)

type TKEClusterConfigSpec struct {
	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"clusterAdvancedSettings,omitempty" yaml:"clusterAdvancedSettings,omitempty"`
	ClusterBasicSettings    *ClusterBasicSettings    `json:"clusterBasicSettings,omitempty" yaml:"clusterBasicSettings,omitempty"`
	ClusterCIDRSettings     *ClusterCIDRSettings     `json:"clusterCIDRSettings,omitempty" yaml:"clusterCIDRSettings,omitempty"`
	ClusterEndpoint         *ClusterEndpoint         `json:"clusterEndpoint,omitempty" yaml:"clusterEndpoint,omitempty"`
	ClusterID               string                   `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	ExtensionAddon          []ExtensionAddon         `json:"extensionAddon,omitempty" yaml:"extensionAddon,omitempty"`
	Imported                bool                     `json:"imported,omitempty" yaml:"imported,omitempty"`
	NodePoolList            []NodePoolDetail         `json:"nodePoolList,omitempty" yaml:"nodePoolList,omitempty"`
	Region                  string                   `json:"region,omitempty" yaml:"region,omitempty"`
	RunInstancesForNode     *RunInstancesForNode     `json:"runInstancesForNode,omitempty" yaml:"runInstancesForNode,omitempty"`
	TKECredentialSecret     string                   `json:"tkeCredentialSecret,omitempty" yaml:"tkeCredentialSecret,omitempty"`
}
