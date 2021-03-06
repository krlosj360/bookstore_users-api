package msql_utils

import (
	"bookstore_users-api/src/util/errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}

		return errors.NewInternalServerError("error parsing msql response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalid data"))
	}
	return errors.NewInternalServerError("error processing request")
}
