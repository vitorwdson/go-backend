package db

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var PgDuplicateError = &pgconn.PgError{Code: "23505"}

func IsErrorCode(err error, pgError *pgconn.PgError) bool {
	if errors.As(err, &pgError) {
		return true
	}

	return false
}
