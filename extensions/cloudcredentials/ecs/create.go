package ecs

import (
	"github.com/rancher/shepherd/clients/rancher"
	management "github.com/rancher/shepherd/clients/rancher/generated/management/v3"
	"github.com/rancher/shepherd/extensions/cloudcredentials"
	"github.com/rancher/shepherd/pkg/config"
)

const ecsCloudCredNameBase = "ecsCloudCredential"

func CreateECSCloudCredentials(rancherClient *rancher.Client) (*cloudcredentials.CloudCredential, error) {
	var aliyunECSCredentialConfig cloudcredentials.AliyunECSCredentialConfig
	config.LoadConfig(cloudcredentials.AliyunECSCredentialConfigurationFileKey, &aliyunECSCredentialConfig)

	cloudCredential := cloudcredentials.CloudCredential{
		Name:                      ecsCloudCredNameBase,
		AliyunECSCredentialConfig: &aliyunECSCredentialConfig,
	}

	resp := &cloudcredentials.CloudCredential{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(management.CloudCredentialType, cloudCredential, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
