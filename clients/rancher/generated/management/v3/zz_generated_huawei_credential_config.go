package client

const (
	HuaweiCredentialConfigType           = "huaweiCredentialConfig"
	HuaweiCredentialConfigFieldAccessKey = "accessKey"
	HuaweiCredentialConfigFieldProjectID = "projectID"
	HuaweiCredentialConfigFieldRegionID  = "regionID"
	HuaweiCredentialConfigFieldSecretKey = "secretKey"
)

type HuaweiCredentialConfig struct {
	AccessKey string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`
	ProjectID string `json:"projectID,omitempty" yaml:"projectID,omitempty"`
	RegionID  string `json:"regionID,omitempty" yaml:"regionID,omitempty"`
	SecretKey string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`
}
