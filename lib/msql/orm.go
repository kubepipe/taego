package msql

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"taego/lib/mlog"
	"taego/lib/mtrace"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type msql struct {
	db *sql.DB
}

func NewSQL(dataSourceName string) SQL {
	db, err := initdb(dataSourceName)
	if err != nil {
		mlog.Errorf("GetSQL err:%v", err)
	}
	return &msql{
		db: db,
	}
}

func (m *msql) Query(ctx context.Context, query string, args ...any) Rows {
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

func (m *msql) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	trace := mtrace.SubTrace(ctx, query)
	defer trace.Done()

	return m.db.ExecContext(ctx, query, args...)
}

func (m *msql) Close() error {
	return m.db.Close()
}

var (
	dbs = make(map[string]*sql.DB)
	mu  sync.RWMutex
)

func initdb(dataSourceName string) (*sql.DB, error) {
	mu.RLock()
	db, ok := dbs[dataSourceName]
	mu.RUnlock()
	if ok {
		return db, nil
	}

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)

	mu.Lock()
	defer mu.Unlock()
	dbs[dataSourceName] = db
	go db.Ping()
	return db, nil
}
