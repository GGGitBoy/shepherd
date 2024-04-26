package client

const (
	TKECredentialConfigType                 = "tkeCredentialConfig"
	TKECredentialConfigFieldAccessKeyID     = "accessKeyId"
	TKECredentialConfigFieldAccessKeySecret = "accessKeySecret"
)

type TKECredentialConfig struct {
	AccessKeyID     string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty" yaml:"accessKeySecret,omitempty"`
}
