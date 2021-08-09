package quectel

import "testing"

func TestClient_GetBillingGroupList(t *testing.T) {
	resp, err := c.GetBillingGroupList()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(resp)
}
