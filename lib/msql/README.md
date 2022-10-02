# msql

1.首先需要初始化一个msql.SQL对象，用来连接db

```go
var user msql.SQL

func init() {
	user = msql.NewSQL("user:password@/dbname")
}
```

这里的```user:password@/dbname```表示DSN: [参考](https://github.com/go-sql-driver/mysql#dsn-data-source-name).

2.然后定义一个struct，表示表结构，以user表为例：

```sql
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

定义的struct如下：

```go
type User struct {
	Id   int64  `db:"id"` // db:"id" means table user have id fields
	Name string `db:"name"`
}
```

其中```db:"id"```表示表结构中对应字段名为id

## insert

```go
id := 1
name := "tom"
if _, err := user.Exec(ctx, "insert into user(id, name) values(?,?)", id, name); err != nil {
	return err
}
```

## delete

```go
id := 1
if _, err := user.Exec(ctx, "delete from user where id=?", id); err != nil {
	return err
}
```

## update

```go
id := 1
name := "liming"
if _, err := user.Exec(ctx, "update user set name=? where id=?", name, id); err != nil {
	return err
}
```

## select

```go
names := []string{}
_ = user.Query(ctx, "select name from user limit 10").Scan(&names)
```

scan方法的参数可传入5种类型：

* 结构体指针 &User{}
* 结构体切片 &[]User{}
* 指针结构体切片 &[]*User{}
* 任何实现了Scanner方法的类型，具体参见 [database/sql#Rows.Scan](https://pkg.go.dev/database/sql#Rows.Scan)
* Scanner类型的切片

以下分别为5种类型的示例：

1.结构体指针

```go
u := User{}
id := 1
if err := user.Query(ctx, "select id,name from user where id=?", id).Scan(&u); err != nil {
	return err
}
```

2.结构体切片

```go
us := []User{}
if err := user.Query(ctx, "select id,name from user").Scan(&us); err != nil {
	return err
}
```

3.指针结构体切片

```go
us := []*User{}
if err := user.Query(ctx, "select id,name from user").Scan(&us); err != nil {
	return err
}
```

4.Scanner类型，包含golang常见的基本数据类型，例如：

```go
*string
*[]byte
*int, *int8, *int16, *int32, *int64
*uint, *uint8, *uint16, *uint32, *uint64
*bool
*float32, *float64
*interface{}
*RawBytes
*Rows (cursor value)
any type implementing Scanner (see Scanner docs)
```

示例：

```go
var name string
id := 1
if err := user.Query(ctx, "select name from user where id=?", id).Scan(&name); err != nil {
	return err
}
```

5.Scanner类型的切片：

```go
var names []string
if err := user.Query(ctx, "select name from user").Scan(&names); err != nil {
	return err
}
```

另外对于SQL接口定义的```Query(context.Context, string, ...any) Rows```，如果没有调用返回的Rows的Scan方法，需要使用Rows.Close释放连接：

```go
rows := user.Query(ctx, "select 1")
defer rows.Close()
```

而Rows.Scan方法内部已经默认调用了Close方法，因此对于调用了Scan的场景，不需要再手动Close.

## 事务

