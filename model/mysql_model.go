package model

// 用户表结构体 结构体字段名自己定义，右边的tag为数据库里面的字段名
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}