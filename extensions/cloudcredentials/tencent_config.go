package cloudcredentials

// The json/yaml config key for the tke cloud credential config
const TencentCredentialConfigurationFileKey = "tencentCredentials"

// TencentCredentialConfig is configuration need to create an tke cloud credential
type TencentCredentialConfig struct {
	AccessKeyID     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
}
