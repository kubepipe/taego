package msql

import (
	"context"
	"database/sql"
)

type SQL interface {
	Query(context.Context, string, ...any) Rows
	Exec(context.Context, string, ...any) (sql.Result, error)

	Close() error
}
