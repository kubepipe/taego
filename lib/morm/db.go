package morm

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"taego/lib/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbs = make(map[string]*sql.DB)
	mu  sync.RWMutex
)

func initdb(dbname string) (*sql.DB, error) {
	mu.RLock()
	db, ok := dbs[dbname]
	mu.RUnlock()
	if ok {
		return db, nil
	}

	dsn := config.Config.UString(fmt.Sprintf("mysql.%s.address", dbname))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)

	mu.Lock()
	defer mu.Unlock()
	dbs[dbname] = db
	go db.Ping()
	return db, nil
}
