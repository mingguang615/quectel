package quectel

import "net/url"

type SendSms struct {
	Response
	SuccessList []struct {
		Msisdn string `json:"msisdn"`
		SmsId  string `json:"sms_id"`
	} `json:"successList"`
	ErrorList []string `json:"errorList"`
}

/*
msisdns:手机号码列表，用英文逗号分隔，最多100个手机号码
content:短信内容
*/
func (c *client) SendSms(msisdns, conten string) (*SendSms, error) {
	param := url.Values{}
	param.Add("msisdns", msisdns)
	param.Add("conten", conten)
	param.Add("method", "fc.function.sms.send")
	result := new(SendSms)
	err := c.post(param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
