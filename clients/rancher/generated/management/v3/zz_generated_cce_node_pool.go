package client

const (
	CCENodePoolType                      = "cceNodePool"
	CCENodePoolFieldAutoscaling          = "autoscaling"
	CCENodePoolFieldCustomSecurityGroups = "customSecurityGroups"
	CCENodePoolFieldID                   = "nodePoolID"
	CCENodePoolFieldInitialNodeCount     = "initialNodeCount"
	CCENodePoolFieldName                 = "name"
	CCENodePoolFieldNodeTemplate         = "nodeTemplate"
	CCENodePoolFieldPodSecurityGroups    = "podSecurityGroups"
	CCENodePoolFieldType                 = "type"
)

type CCENodePool struct {
	Autoscaling          *CCENodePoolNodeAutoscaling `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
	CustomSecurityGroups []string                    `json:"customSecurityGroups,omitempty" yaml:"customSecurityGroups,omitempty"`
	ID                   string                      `json:"nodePoolID,omitempty" yaml:"nodePoolID,omitempty"`
	InitialNodeCount     int64                       `json:"initialNodeCount,omitempty" yaml:"initialNodeCount,omitempty"`
	Name                 string                      `json:"name,omitempty" yaml:"name,omitempty"`
	NodeTemplate         *CCENodeTemplate            `json:"nodeTemplate,omitempty" yaml:"nodeTemplate,omitempty"`
	PodSecurityGroups    []string                    `json:"podSecurityGroups,omitempty" yaml:"podSecurityGroups,omitempty"`
	Type                 string                      `json:"type,omitempty" yaml:"type,omitempty"`
}
