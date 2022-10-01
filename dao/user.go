package dao

import (
	"context"

	"taego/lib/config"
	"taego/lib/morm"
)

/*
create database demo;

USE demo;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

var user morm.ORM

func init() {
	dataSourceName := config.Config.UString("mysql.demo.address")
	user = morm.NewORM(dataSourceName)
}

type User struct {
	Id   int64  `json:"id" db:"id"` // db:"id" means table user have id fields
	Name string `json:"name" db:"name"`
}

// example
func GetUserNames(ctx context.Context) (names []string) {
	_ = user.Query(ctx, "select name from user limit 10").Scan(&names)
	return
}
