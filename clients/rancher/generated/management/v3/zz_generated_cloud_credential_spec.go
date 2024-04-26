package client

const (
	CloudCredentialSpecType                        = "cloudCredentialSpec"
	CloudCredentialSpecFieldDescription            = "description"
	CloudCredentialSpecFieldDisplayName            = "displayName"
	CloudCredentialSpecFieldHuaweiCredentialConfig = "huaweicredentialConfig"
	CloudCredentialSpecFieldS3CredentialConfig     = "s3credentialConfig"
	CloudCredentialSpecFieldTKECredentialConfig    = "tkecredentialConfig"
)

type CloudCredentialSpec struct {
	Description            string                  `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName            string                  `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	HuaweiCredentialConfig *HuaweiCredentialConfig `json:"huaweicredentialConfig,omitempty" yaml:"huaweicredentialConfig,omitempty"`
	S3CredentialConfig     *S3CredentialConfig     `json:"s3credentialConfig,omitempty" yaml:"s3credentialConfig,omitempty"`
	TKECredentialConfig    *TKECredentialConfig    `json:"tkecredentialConfig,omitempty" yaml:"tkecredentialConfig,omitempty"`
}
