package client

const (
	CCENodePoolNodeAutoscalingType                       = "cceNodePoolNodeAutoscaling"
	CCENodePoolNodeAutoscalingFieldEnable                = "enable"
	CCENodePoolNodeAutoscalingFieldMaxNodeCount          = "maxNodeCount"
	CCENodePoolNodeAutoscalingFieldMinNodeCount          = "minNodeCount"
	CCENodePoolNodeAutoscalingFieldPriority              = "priority"
	CCENodePoolNodeAutoscalingFieldScaleDownCooldownTime = "scaleDownCooldownTime"
)

type CCENodePoolNodeAutoscaling struct {
	Enable                bool  `json:"enable,omitempty" yaml:"enable,omitempty"`
	MaxNodeCount          int64 `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`
	MinNodeCount          int64 `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
	Priority              int64 `json:"priority,omitempty" yaml:"priority,omitempty"`
	ScaleDownCooldownTime int64 `json:"scaleDownCooldownTime,omitempty" yaml:"scaleDownCooldownTime,omitempty"`
}
