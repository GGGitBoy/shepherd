package cvm

import (
	"github.com/rancher/shepherd/pkg/config"
	tccommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvmapi "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// Client is a struct that wraps the needed TencentCVMConfig object, and cvm.Client which makes the actual calls to tencent cloud.
type Client struct {
	Client       *cvmapi.Client
	ClientConfig *TencentCVMConfigs
}

// NewClient is a constructor that creates an *CVMClient which a wrapper for a "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm" session and
// the tencent cloud cvm config.
func NewClient() (*Client, error) {
	tencentCVMClientConfig := new(TencentCVMConfigs)

	config.LoadConfig(ConfigurationFileKey, tencentCVMClientConfig)

	credential := tccommon.NewCredential(tencentCVMClientConfig.TencentAccessKeyID, tencentCVMClientConfig.TencentAccessKeySecret)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, err := cvmapi.NewClient(credential, tencentCVMClientConfig.Region, cpf)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client:       client,
		ClientConfig: tencentCVMClientConfig,
	}, nil
}
