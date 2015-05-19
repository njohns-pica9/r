package storage

type Client struct {
	Id          string      `gorethink:"client_id"`
	Secret      string      `gorethink:"client_secret"`
	RedirectUri string      `gorethink:"redirect_uri"`
	UserData    interface{} `gorethink:"user_data"`
}

func (c Client) GetId() string {
	return c.Id
}

func (c Client) GetSecret() string {
	return c.Secret
}

func (c Client) GetRedirectUri() string {
	return c.RedirectUri
}

func (c Client) GetUserData() interface{} {
	return c.UserData
}

func (c Client) FieldMap() map[string]string {
	return map[string]string{
		"Id":          "client_id",
		"Secret":      "client_secret",
		"RedirectUri": "redirect_uri",
		"UserData":    "user_data",
	}
}
