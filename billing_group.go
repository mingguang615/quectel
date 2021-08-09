package quectel

import "net/url"

type BillingGroupList struct {
	Response
	BillingGroups []struct {
		Name        string `json:"name"`        // 计费组名称
		Code        string `json:"code"`        // 编号
		Carrier     string `json:"carrier"`     // 运行商
		CardType    string `json:"cardType"`    // 资产类型
		SetMeal     string `json:"setMeal"`     // 套餐
		ProductCode string `json:"productCode"` // 产品编码
	} `json:"billingGroups"`
}

func (c *client) GetBillingGroupList() (*BillingGroupList, error) {
	param := url.Values{}
	param.Add("method", "fc.function.billing_group.list")
	result := new(BillingGroupList)
	err := c.post(param, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
