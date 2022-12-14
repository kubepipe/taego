package dao

import (
	"taego/lib/config"
	"taego/lib/msql"
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

var user msql.SQL

func init() {
	dataSourceName := config.Config.UString("mysql.demo.address")
	user = msql.NewSQL(dataSourceName)
}

type User struct {
	Id   int64  `json:"id" db:"id"` // db:"id" means table user have id fields
	Name string `json:"name" db:"name"`
}
