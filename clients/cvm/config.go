package cvm

// The json/yaml config key for the TencentCVMonfig
const ConfigurationFileKey = "tencentCVMConfigs"

// TencentCVMConfigs is the TencentCloud authentication configuration for accessing and launching cvm instances
type TencentCVMConfigs struct {
	TencentCVMConfig       []TencentCVMConfig `json:"tencentCVMConfig" yaml:"tencentCVMConfig"`
	TencentAccessKeyID     string             `json:"tencentAccessKeyID" yaml:"tencentAccessKeyID"`
	TencentAccessKeySecret string             `json:"tencentAccessKeySecret" yaml:"tencentAccessKeySecret"`
	Region                 string             `json:"region" yaml:"region"`
}

// TencentCVMConfig is the instance-specific configuration needed to launch cvm instances in TencentCloud
type TencentCVMConfig struct {
	Zone                    string   `json:"zone" yaml:"zone"`
	InstanceType            string   `json:"instanceType" yaml:"instanceType"`
	ImageID                 string   `json:"imageId" yaml:"imageId"`
	InstanceName            string   `json:"instanceName" yaml:"instanceName"`
	KeyIDs                  []string `json:"keyIds" yaml:"keyIds"`
	SecurityGroupIDs        []string `json:"securityGroupIds" yaml:"securityGroupIds"`
	TencentSSHKey           string   `json:"tencentSSHKey" yaml:"tencentSSHKey"`
	LaunchTemplateID        string   `json:"launchTemplateId" yaml:"launchTemplateId"`
	TencentUser             string   `json:"tencentUser" yaml:"tencentUser"`
	InternetChargeType      string   `json:"internetChargeType" yaml:"internetChargeType"`
	VpcID                   string   `json:"vpcId" yaml:"vpcId"`
	SubnetID                string   `json:"subnetId" yaml:"subnetId"`
	PublicIPAssigned        bool     `json:"publicIpAssigned" yaml:"publicIpAssigned"`
	InternetMaxBandwidthOut int      `json:"internetMaxBandwidthOut" yaml:"internetMaxBandwidthOut"`
	UserData                string   `json:"userData" yaml:"userData"`
	Roles                   []string `json:"roles" yaml:"roles"`
}
