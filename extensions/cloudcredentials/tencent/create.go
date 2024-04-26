package tencent

import (
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/cloudcredentials"
	"github.com/rancher/shepherd/pkg/config"
)

const tencentCloudCredNameBase = "tencentCloudCredential"

// CreateTencentCloudCredentials is a helper function that takes the rancher Client as a parameter and creates
// an Tencent cloud credential, and returns the CloudCredential response
func CreateTencentCloudCredentials(rancherClient *rancher.Client) (*cloudcredentials.CloudCredential, error) {
	var tencentCredentialConfig cloudcredentials.TencentCredentialConfig
	config.LoadConfig(cloudcredentials.TencentCredentialConfigurationFileKey, &tencentCredentialConfig)

	cloudCredential := cloudcredentials.CloudCredential{
		Name:                    tencentCloudCredNameBase,
		TencentCredentialConfig: &tencentCredentialConfig,
	}

	resp := &cloudcredentials.CloudCredential{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(management.CloudCredentialType, cloudCredential, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
