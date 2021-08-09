package quectel

import "testing"

const (
	appKey    = "123123"
	appSecret = "BBBBBBBBBB"
)

func TestNewClient(t *testing.T) {
	c := NewClient(appKey, appSecret)
	t.Log(c)
}
