package storage

import "github.com/RangelReale/osin"

type AccessData struct {
	*osin.AccessData
}

func (c AccessData) FieldMap() map[string]string {
	return map[string]string{
		"Client":        "client",
		"AuthorizeData": "authorize_data",
		"AccessData":    "access_data",
		"AccessToken":   "access_token",
		"RefreshToken":  "refresh_token",
		"ExpiresIn":     "expires_in",
		"Scope":         "scope",
		"RedirectUri":   "redirect_uri",
		"CreatedAt":     "created_at",
		"UserData":      "user_data",
	}
}
