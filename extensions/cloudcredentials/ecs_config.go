package cloudcredentials

const AliyunECSCredentialConfigurationFileKey = "ecsCredentials"

type AliyunECSCredentialConfig struct {
	AccessKeyID     string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty" yaml:"accessKeySecret,omitempty"`
}
