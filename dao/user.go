package dao

/*
create database demo;
*/

type demodb struct{}

func (d demodb) GetDBName() string {
	return "demo"
}

/*
USE demo;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type User struct {
	demodb
	Id   int64  `json:"id" db:"id"` // db:"id" means table user have id fields
	Name string `json:"name" db:"name"`
}
