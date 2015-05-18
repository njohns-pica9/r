package storage

import "github.com/RangelReale/osin"

type AuthorizeData struct {
	*osin.AuthorizeData
}

func (c AuthorizeData) FieldMap() map[string]string {
	return map[string]string{
		"Client":      "client",
		"Code":        "code",
		"ExpiresIn":   "expires_in",
		"Scope":       "scope",
		"RedirectUri": "redirect_uri",
		"State":       "state",
		"CreatedAt":   "created_at",
		"UserData":    "user_data",
	}
}
