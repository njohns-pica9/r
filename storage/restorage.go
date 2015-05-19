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
		Address:  "localhost:28015",
		Database: "auth",
		MaxIdle:  10,
		MaxOpen:  10,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &Restorage{
		session: session,
	}
}

func (s *Restorage) Reconnect() {
	err := s.session.Reconnect()

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Restorage) Clone() osin.Storage {
	return s
}

func (s *Restorage) Close() {
	err := s.session.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func (s *Restorage) GetClient(id string) (osin.Client, error) {
	s.Reconnect()

	var client Client

	cursor, err := r.Db("auth").Table("clients").
		GetAllByIndex("client_id", id).
		Limit(1).
		Run(s.session)

	log.Printf("Session: %v\n", s.session)

	if err != nil {
		log.Println(err)
		return client, err
	}

	cursor.Next(&client)

	log.Printf("GetClient: %v\n", client)

	return client, nil
}

func (s *Restorage) SaveAuthorize(data *osin.AuthorizeData) error {
	s.Reconnect()

	redata := convertToAuthorizeData(data)

	_, err := r.Db("auth").Table("client_authorizations").Insert(redata).RunWrite(s.session)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("SaveAuthorize: %v\n", data)

	return nil
}

func (s *Restorage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	s.Reconnect()

	var authroize_data AuthorizeData

	cursor, err := r.Db("auth").Table("client_authorizations").
		GetAllByIndex("code", code).
		Limit(1).
		Run(s.session)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cursor.Next(&authroize_data)

	log.Printf("LoadAuthorize: %v\n", authroize_data)

	return convertFromAuthorizeData(&authroize_data), nil
}

func (s *Restorage) RemoveAuthorize(code string) error {
	s.Reconnect()

	_, err := r.Db("auth").Table("client_authorizations").
		GetAllByIndex("code", code).
		Limit(1).
		Delete().
		RunWrite(s.session)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("RemoveAuthorize: %v\n", code)

	return nil
}

func (s *Restorage) SaveAccess(data *osin.AccessData) error {
	s.Reconnect()

	redata := convertToAccessData(data)

	_, err := r.Db("auth").Table("client_access").Insert(redata).RunWrite(s.session)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("SaveAccess: %v\n", data)

	return nil
}

func (s *Restorage) LoadAccess(code string) (*osin.AccessData, error) {
	s.Reconnect()

	var access_data AccessData

	cursor, err := r.Db("auth").Table("client_access").
		GetAllByIndex("access_token", code).
		Limit(1).
		Run(s.session)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cursor.Next(&access_data)

	log.Printf("LoadAccess: %v\n", access_data)

	return convertFromAccessData(&access_data), nil
}

func (s *Restorage) RemoveAccess(code string) error {
	s.Reconnect()

	_, err := r.Db("auth").Table("client_access").
		GetAllByIndex("access_token", code).
		Limit(1).
		Delete().
		RunWrite(s.session)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("RemoveAccess: %v\n", code)

	return nil
}

func (s *Restorage) LoadRefresh(code string) (*osin.AccessData, error) {
	s.Reconnect()

	var access_data osin.AccessData

	cursor, err := r.Db("auth").Table("client_access").
		GetAllByIndex("refresh_token", code).
		Limit(1).
		Run(s.session)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	cursor.Next(&access_data)

	return &access_data, nil
}

func (s *Restorage) RemoveRefresh(code string) error {
	s.Reconnect()

	_, err := r.Db("auth").Table("client_access").
		GetAllByIndex("refresh_token", code).
		Limit(1).
		Delete().
		RunWrite(s.session)

	if err != nil {
		return err
	}

	return nil
}
