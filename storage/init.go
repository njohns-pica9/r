package storage

import (
	r "github.com/dancannon/gorethink"
	"log"
)

func (s *Restorage) Init() {
	s.initDb()
	s.initTables()
	s.initIndexes()
}

func (s *Restorage) initDb() {
	_, err := r.DbCreate("auth").RunWrite(s.session)
	if err != nil {
		log.Println(err)
	}
}

func (s *Restorage) initTables() {
	tables := []string{
		"clients",
		"client_authorizations",
		"client_access",
	}

	for _, v := range tables {
		_, err := r.Db("auth").TableCreate(v).RunWrite(s.session)
		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Restorage) initIndexes() {
	indexes := map[string]string{
		"client_id": "clients",
		"code": "client_authorizations",
		"access_token": "client_access",
		"refresh_token": "client_access",
	}

	for index, table := range indexes {
		_, err := r.Db("auth").Table(table).IndexCreate(index).RunWrite(s.session)
		if err != nil {
			log.Println(err)
		}
	}
}