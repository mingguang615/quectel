package quectel

import "testing"

func TestClient_GetBillingGroupList(t *testing.T) {
	resp, err := c.GetBillingGroupList()
	if err != nil {
		panic(err)
	}
	t.Log(resp)
}
