package tencent

import (
	"context"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/sms"

	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	smsSDK "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// 当前API只支持发送中国区域的手机号，也就是+86
type tencentSms struct {
	client        *smsSDK.Client
	applicationId string
	signName      string
	templateId    string
}

func NewSms(c *instance.TencentSmsConfig) sms.Sms {
	client, err := instance.NewTencentSms(c)
	if err != nil {
		panic(errors.Errorf("init tencent sms client failed! err: %v", err))
	}
	return &tencentSms{
		client:        client,
		applicationId: c.ApplicationId,
		signName:      c.SignName,
		templateId:    c.TemplateId,
	}
}

func (t *tencentSms) buildBasicRequest() *smsSDK.SendSmsRequest {
	request := smsSDK.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(t.applicationId)
	request.SignName = common.StringPtr(t.signName)
	request.TemplateId = common.StringPtr(t.templateId)
	return request
}

func (t *tencentSms) SendToSingle(c context.Context, phone string, paramSet []string) error {
	request := t.buildBasicRequest()
	request.TemplateParamSet = common.StringPtrs(paramSet)
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone})

	_, err := t.client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}

func (t *tencentSms) SendToMultiple(c context.Context, phones []string, paramSet []string) error {
	if len(phones) > 200 {
		return errors.New("手机号数量超过200")
	}
	request := t.buildBasicRequest()
	request.TemplateParamSet = common.StringPtrs(paramSet)
	for i := range phones {
		phones[i] = "+86" + phones[i]
	}
	request.PhoneNumberSet = common.StringPtrs(phones)

	_, err := t.client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}
