package quectel

const (
	api = "http://api.quectel.com/openapi/router"
)

type client struct {
	appKey    string
	appSecret string
}

func NewClient(key, secret string) *client {
	return &client{appKey: key, appSecret: secret}
}
