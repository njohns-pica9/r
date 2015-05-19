package storage

import (
	"github.com/RangelReale/osin"
	"time"
)

type AccessData struct {
	Client        Client         `gorethink:"client"`
	AuthorizeData *AuthorizeData `gorethink:"authorize_data"`
	AccessData    *AccessData    `gorethink:"access_data"`
	AccessToken   string         `gorethink:"access_token"`
	RefreshToken  string         `gorethink:"refresh_token"`
	ExpiresIn     int32          `gorethink:"expires_in"`
	Scope         string         `gorethink:"scope"`
	RedirectUri   string         `gorethink:"redirect_uri"`
	CreatedAt     time.Time      `gorethink:"created_at"`
	UserData      interface{}    `gorethink:"user_data"`
}

func convertToAccessData(data *osin.AccessData) *AccessData {
	if data == nil {
		return nil
	}

	return &AccessData{
		Client:        data.Client.(Client),
		AuthorizeData: convertToAuthorizeData(data.AuthorizeData),
		AccessData:    convertToAccessData(data.AccessData),
		AccessToken:   data.AccessToken,
		RefreshToken:  data.RefreshToken,
		ExpiresIn:     data.ExpiresIn,
		Scope:         data.Scope,
		RedirectUri:   data.RedirectUri,
		CreatedAt:     data.CreatedAt,
		UserData:      data.UserData,
	}
}

func convertFromAccessData(data *AccessData) *osin.AccessData {
	if data == nil {
		return nil
	}

	return &osin.AccessData{
		Client:        data.Client,
		AuthorizeData: convertFromAuthorizeData(data.AuthorizeData),
		AccessData:    convertFromAccessData(data.AccessData),
		AccessToken:   data.AccessToken,
		RefreshToken:  data.RefreshToken,
		ExpiresIn:     data.ExpiresIn,
		Scope:         data.Scope,
		RedirectUri:   data.RedirectUri,
		CreatedAt:     data.CreatedAt,
		UserData:      data.UserData,
	}
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
