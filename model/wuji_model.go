package model

//无极测试表数据类型
type WujiData struct {
	ID   	int64  `wuji:"id"` //主键
	Name 	string	`wuji:"name"`
}
