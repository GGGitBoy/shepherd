package client

const (
	AutoScalingGroupParaType                      = "autoScalingGroupPara"
	AutoScalingGroupParaFieldAutoScalingGroupName = "autoScalingGroupName"
	AutoScalingGroupParaFieldDesiredCapacity      = "desiredCapacity"
	AutoScalingGroupParaFieldMaxSize              = "maxSize"
	AutoScalingGroupParaFieldMinSize              = "minSize"
	AutoScalingGroupParaFieldSubnetIDs            = "subnetIds"
	AutoScalingGroupParaFieldVpcID                = "vpcId"
)

type AutoScalingGroupPara struct {
	AutoScalingGroupName string   `json:"autoScalingGroupName,omitempty" yaml:"autoScalingGroupName,omitempty"`
	DesiredCapacity      int64    `json:"desiredCapacity,omitempty" yaml:"desiredCapacity,omitempty"`
	MaxSize              int64    `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`
	MinSize              int64    `json:"minSize,omitempty" yaml:"minSize,omitempty"`
	SubnetIDs            []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
	VpcID                string   `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
