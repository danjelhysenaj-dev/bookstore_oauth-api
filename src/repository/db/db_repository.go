package db

import (
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/clients/cassandra"
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/domain/access_token"
	"github.com/danjelhysenaj-dev/bookstore_auth-api/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
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
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
