package quectel

import "testing"

func TestClient_SendSms(t *testing.T) {
	resp, err := c.SendSms("123,321", "test")
	if err != nil {
		panic(err)
	}
	t.Log(resp)
}
