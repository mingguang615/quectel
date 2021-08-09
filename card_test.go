package quectel

import "testing"

var c = NewClient(appKey, appSecret)

func TestClient_GetCardInfo(t *testing.T) {
	resp, err := c.GetCardInfo("898604781021111071500")
	t.Log(resp, err)
}

func TestClient_GetCardList(t *testing.T) {
	resp, err := c.GetCardList(1,100)
	t.Log(resp, err)
}

func TestClient_GetCardRealTimeStatus(t *testing.T) {
	resp, err := c.GetCardRealTimeStatus("898604781020500222200")
	t.Log(resp, err)
}
