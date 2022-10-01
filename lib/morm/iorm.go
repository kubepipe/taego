package morm

import (
	"context"
	"database/sql"
)

type ORM interface {
	Query(context.Context, string, ...any) Rows
	Exec(context.Context, string, ...any) (sql.Result, error)

	Close() error
}
