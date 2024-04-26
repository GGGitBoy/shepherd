package huawei

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	huawei_ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	"github.com/rancher/shepherd/pkg/config"
)

// Client is a struct that wraps the needed ECSConfig object
type Client struct {
	Client       *huawei_ecs.EcsClient
	ClientConfig *ECSConfigs
}

// NewClient is a constructor that creates an *Client.
func NewClient() (*Client, error) {
	huaweiClientConfig := new(ECSConfigs)

	config.LoadConfig(ConfigurationFileKey, huaweiClientConfig)

	auth := basic.NewCredentialsBuilder().
		WithAk(huaweiClientConfig.HuaweiAccessKey).
		WithSk(huaweiClientConfig.HuaweiSecretKey).
		Build()
	client := huawei_ecs.NewEcsClient(
		huawei_ecs.EcsClientBuilder().
			WithRegion(region.ValueOf(huaweiClientConfig.Region)).
			WithCredential(auth).
			Build())

	return &Client{
		Client:       client,
		ClientConfig: huaweiClientConfig,
	}, nil
}
