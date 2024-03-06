package db

import (
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/clients/cassandra"
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/domain/access_token"
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	//TODO: Implement get access token from CassandraDB.
	_, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
