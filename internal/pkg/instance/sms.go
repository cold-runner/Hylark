package instance

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	smsSDK "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type TencentSmsConfig struct {
	SecretId      string `mapstructure:"secret-id"`
	SecretKey     string `mapstructure:"secret-key"`
	ApplicationId string `mapstructure:"application-id"`
	SignName      string `mapstructure:"sign-name"`
	TemplateId    string `mapstructure:"template-id"`
}

func NewTencentSms(c *TencentSmsConfig) (*smsSDK.Client, error) {
	credential := common.NewCredential(c.SecretId, c.SecretKey)
	client, err := smsSDK.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	if err != nil {
		return nil, err
	}
	return client, nil
}
