package quectel

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	ResultCode   int64  `json:"resultCode"`
	ErrorMessage string `json:"errorMessage"`
}

func (c *client) getUrl(param url.Values) string {
	return fmt.Sprintf("%s?%s", api, param.Encode())
}

func (c *client) post(param url.Values, result interface{}) (error) {
	if param == nil {
		param = url.Values{}
	}
	param.Add("appKey", c.appKey)
	param.Add("t", fmt.Sprintf("%d", time.Now().Unix()))
	param.Add("sign", c.sign(param))

	resp, err := http.Post(c.getUrl(param), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//default 5MB
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, int64(5<<20)))
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return json.Unmarshal(body, result)
}