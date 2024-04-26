package huawei

// The json/yaml config key for the ECSConfigs
const ConfigurationFileKey = "huaweiECSConfigs"

// ECSConfigs is the Huawei Cloud authentication configuration for accessing and launching ECS instances
type ECSConfigs struct {
	HuaweiConfig    []ECSConfig `json:"ecsConfig" yaml:"ecsConfig"`
	HuaweiAccessKey string      `json:"huaweiAccessKey" yaml:"huaweiAccessKey"`
	HuaweiSecretKey string      `json:"huaweiSecretKey" yaml:"huaweiSecretKey"`
	HuaweiProjectID string      `json:"huaweiProjectID" yaml:"huaweiProjectID"`
	Region          string      `json:"region" yaml:"region"`
}

// ECSConfig is the instance-specific configuration needed to launch  instances in Huawei Cloud
type ECSConfig struct {
	ImageRef            string   `json:"imageRef" yaml:"imageRef"`
	FlavorRef           string   `json:"flavorRef" yaml:"flavorRef"`
	SSHUser             string   `json:"sshUser" yaml:"sshUser"`
	KeyName             string   `json:"keyName" yaml:"keyName"`
	VpcID               string   `json:"vpcID" yaml:"vpcID"`
	SubnetID            string   `json:"subnetID" yaml:"subnetID"`
	EIPType             string   `json:"eipType" yaml:"eipType"`
	BandwidthSize       int32    `json:"bandwidthSize" yaml:"bandwidthSize"`
	BandwidthShareType  string   `json:"bandwidthShareType" yaml:"bandwidthShareType"`
	BandwidthChargeMode string   `json:"bandwidthChargeMode" yaml:"bandwidthChargeMode"`
	RootVolumeSize      int32    `json:"rootVolumeSize" yaml:"rootVolumeSize"`
	RootVolumeType      string   `json:"rootVolumeType" yaml:"rootVolumeType"`
	DataVolumeSize      int32    `json:"dataVolumeSize" yaml:"dataVolumeSize"`
	DataVolumeType      string   `json:"dataVolumeType" yaml:"dataVolumeType"`
	SecurityGroups      []string `json:"securityGroups" yaml:"securityGroups"`
	UserData            string   `json:"userData" yaml:"userData"`
	Roles               []string `json:"roles" yaml:"roles"`
}
