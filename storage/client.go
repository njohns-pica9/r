package storage

import "github.com/RangelReale/osin"

type Client struct {
	osin.Client
}

func (c Client) FieldMap() map[string]string {
	return map[string]string{
		"Id":          "client_id",
		"Secret":      "client_secret",
		"RedirectUri": "redirect_uri",
		"UserData":    "user_data",
	}
}
