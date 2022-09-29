package morm

import (
	"context"
	"database/sql"

	"taego/lib/mlog"
	"taego/lib/mtrace"

	"go.uber.org/zap"
)

type ORM interface {
	Query(context.Context, string, ...any) Rows
	Exec(context.Context, string, ...any) (sql.Result, error)

	Close() error
}

type morm struct {
	db *sql.DB
}

func GetORM(t Table) ORM {
	db, err := initdb(t.GetDBName())
	if err != nil {
		mlog.Errorf("GetORM err:%v", err)
	}
	return &morm{
		db: db,
	}
}

func (m *morm) Query(ctx context.Context, query string, args ...any) Rows {
	mrows := &mrows{}

	trace := mtrace.SubTrace(ctx, query)
	mrows.trace = trace
	defer func() {
		if mrows.err != nil {
			trace.Log("db query err", zap.Error(mrows.err))
			mrows.Close()
		}
	}()

	srows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		mrows.err = err
		return mrows
	}
	mrows.rows = srows

	return mrows
}

func (m *morm) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	trace := mtrace.SubTrace(ctx, query)
	defer trace.Done()

	return m.db.ExecContext(ctx, query, args...)
}

func (m *morm) Close() error {
	return m.db.Close()
}
