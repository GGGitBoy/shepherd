package client

const (
	NodePoolDetailType                      = "nodePoolDetail"
	NodePoolDetailFieldAutoScalingGroupPara = "autoScalingGroupPara"
	NodePoolDetailFieldClusterID            = "clusterId"
	NodePoolDetailFieldDeletionProtection   = "deletionProtection"
	NodePoolDetailFieldEnableAutoscale      = "enableAutoscale"
	NodePoolDetailFieldLabels               = "labels"
	NodePoolDetailFieldLaunchConfigurePara  = "launchConfigurePara"
	NodePoolDetailFieldName                 = "name"
	NodePoolDetailFieldNodePoolID           = "nodePoolId"
	NodePoolDetailFieldNodePoolOs           = "nodePoolOs"
	NodePoolDetailFieldOsCustomizeType      = "osCustomizeType"
	NodePoolDetailFieldTags                 = "tags"
	NodePoolDetailFieldTaints               = "taints"
)

type NodePoolDetail struct {
	AutoScalingGroupPara *AutoScalingGroupPara `json:"autoScalingGroupPara,omitempty" yaml:"autoScalingGroupPara,omitempty"`
	ClusterID            string                `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	DeletionProtection   bool                  `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`
	EnableAutoscale      bool                  `json:"enableAutoscale,omitempty" yaml:"enableAutoscale,omitempty"`
	Labels               []string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	LaunchConfigurePara  *LaunchConfigurePara  `json:"launchConfigurePara,omitempty" yaml:"launchConfigurePara,omitempty"`
	Name                 string                `json:"name,omitempty" yaml:"name,omitempty"`
	NodePoolID           string                `json:"nodePoolId,omitempty" yaml:"nodePoolId,omitempty"`
	NodePoolOs           string                `json:"nodePoolOs,omitempty" yaml:"nodePoolOs,omitempty"`
	OsCustomizeType      string                `json:"osCustomizeType,omitempty" yaml:"osCustomizeType,omitempty"`
	Tags                 []string              `json:"tags,omitempty" yaml:"tags,omitempty"`
	Taints               []string              `json:"taints,omitempty" yaml:"taints,omitempty"`
}
