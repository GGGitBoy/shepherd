package huawei

import (
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/cloudcredentials"
	"github.com/rancher/shepherd/pkg/config"
)

const huaweiCloudCredNameBase = "huaweiCloudCredential"

// CreateHuaweiCloudCredentials is a helper function that takes the rancher Client as a parameter and creates
// an Huawei cloud credential, and returns the CloudCredential response
func CreateHuaweiCloudCredentials(rancherClient *rancher.Client) (*cloudcredentials.CloudCredential, error) {
	var huaweiCredentialConfig cloudcredentials.HuaweiCredentialConfig
	config.LoadConfig(cloudcredentials.HuaweiCredentialConfigurationFileKey, &huaweiCredentialConfig)

	cloudCredential := cloudcredentials.CloudCredential{
		Name:                   huaweiCloudCredNameBase,
		HuaweiCredentialConfig: &huaweiCredentialConfig,
	}

	resp := &cloudcredentials.CloudCredential{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(management.CloudCredentialType, cloudCredential, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
