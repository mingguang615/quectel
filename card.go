package quectel

import (
	"fmt"
	"net/url"
)

type CardInfo struct {
	Response
	Msisdn       string `json:"msisdn"`
	Iccid        string `json:"iccid"`
	Imsi         string `json:"imsi"`
	Supplier     string `json:"supplier"`     // 运营商
	SetMeal      string `json:"setMeal"`      // 套餐
	BillingCode  string `json:"billingCode"`  // 运营商
	ProductCode  string `json:"productCode"`  // 运营商
	NetworkType  string `json:"networkType"`  // 网络类型
	Active       string `json:"active"`       // 激活状态
	ActivateTime string `json:"activateTime"` // 激活日期
	ExpiryDate   string `json:"expiryDate"`   // 计费结束日期
	/*正常 单项停机 停机 预销号 销号 过户 休眠 待激活 未知*/
	Status      string  `json:"status"`      // 状态
	Flow        float64 `json:"flow"`        // 当月已用流量
	ResidueFlow float64 `json:"residueFlow"` // 剩余流量
	CardType    string  `json:"cardType"`    // 卡类型
}

type CardList struct {
	Response
	Cards []struct {
		Msisdn     string  `json:"msisdn"`
		Iccid      string  `json:"iccid"`
		Carrier    string  `json:"carrier"`     //运营商
		SetMeal    string  `json:"setMeal"`     //套餐
		Active     string  `json:"active"`      // 激活状态
		ActiveDate string  `json:"activeDate"`  //激活日期
		ExpiryDate string  `json:"expiry_date"` // 计费结束日期
		Status     string  `json:"status"`      // 状态
		Flow       float64 `json:"flow"`        //当月已用流量
	}
	Page struct {
		Page       int64 `json:"page"`
		TotalCount int64 `json:"totalCount"`
		PageSize   int64 `json:"pageSize"`
		TotalPage  int64 `json:"totalPage"`
	} `json:"page"`
}

type CardRealTimeStatus struct {
	Response
	RealTimeStatus struct {
		IsSuccess    bool    `json:"issuccess"`
		OnlineStatus string  `json:"onlinestatus"`
		UsedFlow     string  `json:"usedFlow"`
		SurpFlow     string  `json:"surpFlow"`
		TotalFlow    string  `json:"totalFlow"`
		GprsStatus   string  `json:"gprsStatus"`
		Ip           string  `json:"ip"`
		UsedFlowb    float64 `json:"usedFlowb"`
		SurpFlowb    float64 `json:"surpFlowb"`
		TotalFlowb   float64 `json:"totalFlowb"`
		Apn          string  `json:"apn"`
	} `json:"realTimeStatus"`
}

/*
ICCID or MSISDN
*/
func (c *client) GetCardInfo(id string) (*CardInfo, error) {
	param := url.Values{}
	if len(id) == 20 {
		param.Add("iccid", id)
	} else {
		param.Add("msisdn", id)
	}
	param.Add("method", "fc.function.card.info")

	result := new(CardInfo)
	err := c.post(param, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
pageNo:请求的第几页， 默认为第1页
pageSize:每页显示数量， 默认每页显示10张
*/
func (c *client) GetCardList(pageNo, pageSize int64) (*CardList, error) {
	param := url.Values{}
	param.Add("pageNo", fmt.Sprintf("%d", pageNo))
	param.Add("pageSize", fmt.Sprintf("%d", pageSize))
	param.Add("method", "fc.function.card.list")

	result := new(CardList)
	err := c.post(param, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
ICCID or MSISDN
*/
func (c *client) GetCardRealTimeStatus(id string) (*CardRealTimeStatus, error) {
	param := url.Values{}
	if len(id) == 20 {
		param.Add("iccid", id)
	} else {
		param.Add("msisdn", id)
	}
	param.Add("method", "fc.function.card.realtimestatus")

	result := new(CardRealTimeStatus)
	err := c.post(param, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
