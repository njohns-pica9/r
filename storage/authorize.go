package storage

import (
	"github.com/RangelReale/osin"
	"time"
)

type AuthorizeData struct {
	Client      Client      `gorethink:"client"`
	Code        string      `gorethink:"code"`
	ExpiresIn   int32       `gorethink:"expires_in"`
	Scope       string      `gorethink:"scope"`
	RedirectUri string      `gorethink:"redirect_uri"`
	State       string      `gorethink:"state"`
	CreatedAt   time.Time   `gorethink:"created_at"`
	UserData    interface{} `gorethink:"user_data"`
}

func convertToAuthorizeData(data *osin.AuthorizeData) *AuthorizeData {
	if data == nil {
		return nil
	}

	return &AuthorizeData{
		Client:      data.Client.(Client),
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectUri,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData:    data.UserData,
	}
}

func convertFromAuthorizeData(data *AuthorizeData) *osin.AuthorizeData {

	if data == nil {
		return nil
	}

	return &osin.AuthorizeData{
		Client:      data.Client,
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectUri,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData:    data.UserData,
	}
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
