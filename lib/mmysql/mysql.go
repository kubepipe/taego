package mmysql

import (
	"database/sql"
	"log"
	"taego/lib/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", config.Config.UString("mysql.address"))
	if err != nil {
		log.Fatalf("sql.Open failed. %v", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)
}

func Ping() error {
	return db.Ping()
}

//func Demo() {
//	conn, err := db.Conn(context.Background())
//	if err != nil {
//		mlog.Error(err)
//		return
//	}
//
//}
