package sms

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
)

const name = "sms"

type Instance struct {
	RegionId        string `yaml:"region_id"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	SignName        string `yaml:"sign_name"`
	TemplateCode    string `yaml:"template_code"`
	client          *dysmsapi.Client
}

var instance *Instance

func (i *Instance) Load() error {
	instance = i
	core.GetComponentConfiguration(name, i)
	var err error
	i.client, err = dysmsapi.NewClientWithAccessKey(i.RegionId, i.AccessKeyId, i.AccessKeySecret)
	if err != nil {
		log.Logger("component", name).Fatal().Err(err).Send()
		return err
	}
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Send(phone string, code string) error {

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = instance.SignName
	request.TemplateCode = instance.TemplateCode
	request.TemplateParam = fmt.Sprintf(`{code:"%s"}`, code)

	resp, err := instance.client.SendSms(request)
	if err != nil {
		return errors.New(fmt.Sprintf("发送短信失败 %v", err))
	} else if resp.Code != "OK" {
		return errors.New(fmt.Sprintf("发送短信失败 %s", resp.BaseResponse.GetHttpContentString()))
	}

	return nil
}
