package storage

import (
	"github.com/RangelReale/osin"
	r "github.com/dancannon/gorethink"
	"log"
)

type Restorage struct {
	session *r.Session
}

type reclient struct {
	Id          string      `gorethink:"client_id"`
	Secret      string      `gorethink:"client_secret"`
	RedirectUri string      `gorethink:"redirect_uri"`
	UserData    interface{} `gorethink:"user_data"`
}

func NewStorage() *Restorage {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &Restorage{
		session: session,
	}
}

func (s *Restorage) Clone() Restorage {
	return *s
}

func (s *Restorage) Close() {
	err := s.session.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Restorage) GetClient(id string) (Client, error) {
	var client Client

	cursor, err := r.Table("clients").
		GetAllByIndex("client_id", id).
		Limit(1).
		Run(s.session)

	if err != nil {
		return client, err
	}

	cursor.Next(&client)

	return client, nil
}

func (s *Restorage) SaveAuthorize(data *osin.AuthorizeData) error {
	_, err := r.Table("client_authorizations").Insert(data).RunWrite(s.session)

	if err != nil {
		return err
	}

	return nil
}

func (s *Restorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	var authroize_data osin.AuthorizeData

	cursor, err := r.Table("client_authorizations").
		GetAllByIndex("code", code).
		Limit(1).
		Run(s.session)

	if err != nil {
		return nil, err
	}

	cursor.Next(&authroize_data)

	return &authroize_data, nil
}

func (s *Restorage) RemoveAuthorize(code string) error {
	_, err := r.Table("client_authorizations").
		GetAllByIndex("code", code).
		Limit(1).
		Delete().
		RunWrite(s.session)

	if err != nil {
		return err
	}

	return nil
}

func (s *Restorage) SaveAccess(data *osin.AccessData) error {
	_, err := r.Table("client_access").Insert(data).RunWrite(s.session)

	if err != nil {
		return err
	}

	return nil
}

func (s *Restorage) LoadAccess(code string) (*osin.AccessData, error) {
	var access_data osin.AccessData

	cursor, err := r.Table("client_access").
		GetAllByIndex("access_token", code).
		Limit(1).
		Run(s.session)

	if err != nil {
		return nil, err
	}

	cursor.Next(&access_data)

	return &access_data, nil
}

func (s *Restorage) RemoveAccess(code string) error {
	_, err := r.Table("client_access").
		GetAllByIndex("access_token", code).
		Limit(1).
		Delete().
		RunWrite(s.session)

	if err != nil {
		return err
	}

	return nil
}

func (s *Restorage) LoadRefresh(code string) (*osin.AccessData, error) {
	var access_data osin.AccessData

	cursor, err := r.Table("client_access").
		GetAllByIndex("refresh_token", code).
		Limit(1).
		Run(s.session)

	if err != nil {
		return nil, err
	}

	cursor.Next(&access_data)

	return &access_data, nil
}

func (s *Restorage) RemoveRefresh(code string) error {
	_, err := r.Table("client_access").
		GetAllByIndex("refresh_token", code).
		Limit(1).
		Delete().
		RunWrite(s.session)

	if err != nil {
		return err
	}

	return nil
}
