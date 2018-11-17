package ginmysqlstore

import (
	"database/sql"

	"github.com/gin-contrib/sessions"
	gsessions "github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
)

//MySQLStore implement gin-contrib/sessions/store
type MySQLStore struct {
	*mysqlstore.MySQLStore
}

func NewMySQLStore(endpoint string, tableName string, path string, maxAge int, keyPairs ...[]byte) (*MySQLStore, error) {
	s, err := mysqlstore.NewMySQLStore(endpoint, tableName, path, maxAge, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &MySQLStore{s}, nil
}

func NewMySQLStoreFromConnection(db *sql.DB, tableName string, path string, maxAge int, keyPairs ...[]byte) (*MySQLStore, error) {
	s, err := mysqlstore.NewMySQLStoreFromConnection(db, tableName, path, maxAge, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &MySQLStore{s}, nil
}

func (s *MySQLStore) Options(options sessions.Options) {
	if options.MaxAge <= 0 {
		options.MaxAge = s.MySQLStore.Options.MaxAge
	}
	s.MySQLStore.Options = &gsessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
