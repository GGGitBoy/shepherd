package client

const (
	LaunchConfigureParaType                         = "launchConfigurePara"
	LaunchConfigureParaFieldDataDisks               = "dataDisks"
	LaunchConfigureParaFieldInstanceChargeType      = "instanceChargeType"
	LaunchConfigureParaFieldInstanceType            = "instanceType"
	LaunchConfigureParaFieldInternetChargeType      = "internetChargeType"
	LaunchConfigureParaFieldInternetMaxBandwidthOut = "internetMaxBandwidthOut"
	LaunchConfigureParaFieldKeyIDs                  = "keyIds"
	LaunchConfigureParaFieldLaunchConfigurationName = "launchConfigurationName"
	LaunchConfigureParaFieldPublicIpAssigned        = "publicIpAssigned"
	LaunchConfigureParaFieldSecurityGroupIDs        = "securityGroupIds"
	LaunchConfigureParaFieldSystemDisk              = "systemDisk"
)

type LaunchConfigurePara struct {
	DataDisks               []DataDisk `json:"dataDisks,omitempty" yaml:"dataDisks,omitempty"`
	InstanceChargeType      string     `json:"instanceChargeType,omitempty" yaml:"instanceChargeType,omitempty"`
	InstanceType            string     `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
	InternetChargeType      string     `json:"internetChargeType,omitempty" yaml:"internetChargeType,omitempty"`
	InternetMaxBandwidthOut int64      `json:"internetMaxBandwidthOut,omitempty" yaml:"internetMaxBandwidthOut,omitempty"`
	KeyIDs                  []string   `json:"keyIds,omitempty" yaml:"keyIds,omitempty"`
	LaunchConfigurationName string     `json:"launchConfigurationName,omitempty" yaml:"launchConfigurationName,omitempty"`
	PublicIpAssigned        bool       `json:"publicIpAssigned,omitempty" yaml:"publicIpAssigned,omitempty"`
	SecurityGroupIDs        []string   `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`
	SystemDisk              *DataDisk  `json:"systemDisk,omitempty" yaml:"systemDisk,omitempty"`
}
