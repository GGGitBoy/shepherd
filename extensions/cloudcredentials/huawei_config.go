package cloudcredentials

// The json/yaml config key for the aws cloud credential config
const HuaweiCredentialConfigurationFileKey = "huaweiCredentials"

// HuaweiCredentialConfig is configuration need to create an aws cloud credential
type HuaweiCredentialConfig struct {
	AccessKey string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
	ProjectID string `json:"projectID,omitempty" yaml:"projectID,omitempty"`
	RegionID  string `json:"regionID,omitempty" yaml:"regionID,omitempty"`
}
