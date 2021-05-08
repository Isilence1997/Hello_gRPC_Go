package dao

import (
	"context"
	"database/sql"
	"fmt"
	"git.code.oa.com/trpc-go/trpc-database/mysql"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
)

const (
	// 默认测试表名
	defaultTableName = "hello_users"
)

var (
	// mysql 客户端代理
	mysqlClientProxy mysql.Client
)

//初始化mysql
func initMysqlProxy() error{
	mysqlClientProxy = mysql.NewClientProxy(
		"trpc.mysql.mysql.mysql",
		//dsn://user:passwd@tcp(vip:port)/db?timeout=1s&parseTime=true&interpolateParams=true")  mdb使用域名多实例需要加上 &interpolateParams=true
		client.WithTarget("dsn://readuser:uFW0Q_47thjeWM@tcp(shortvideotest.mdb.mig:17073)/mysql?timeout=1s&parseTime=true&interpolateParams=true"),
	)
	// 测试是否连接成功
	var res int
	next := func(rows *sql.Rows) error {
		err := rows.Scan(&res)
		if err != nil {
			return err
		}
		return nil
	}
	err := mysqlClientProxy.Query(context.Background(), next, "SELECT 1")//增加临时列，列值为1
	return err
}
//create
func AcessMysqlInit(ctx context.Context) (rsp string,err error) {
	err = initMysqlProxy()
	if err != nil {
		err = fmt.Errorf("InitMDBClientProxy exec sql `SELECT 1` error, err:%v", err)
		return "", err
	}
	rsp = "connect to mysql successfully!"

	//创建表
	sqlStr := fmt.Sprintf("create table %s (`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(30),`age` int(11),PRIMARY KEY (`id`));", defaultTableName)
	result, err := mysqlClientProxy.Exec(ctx, sqlStr)
	if err != nil {
		err = fmt.Errorf("create table error, err:%+v", err)
		return rsp,err
	}
	LastInsertId,_ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected,_ :=result.RowsAffected()  // 被影响的行数
	rsp += fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ",LastInsertId,RowsAffected)

	return rsp, nil
}
//insert
func AcessMysqlInsert(ctx context.Context) (rsp string,err error) {
	sqlStr := fmt.Sprintf("insert into %s(name,age) values (?,?),(?,?)", defaultTableName)
	result, err := mysqlClientProxy.Exec(ctx, sqlStr, "小明",17, "小红", 18)//插入（小明，3,80）（小红，1,75）
	if err != nil {
		err = fmt.Errorf("insert error, err:%+v", err)
		return "",err
	}
	LastInsertId,_ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected,_ :=result.RowsAffected()  // 被影响的行数
	rsp = fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ",LastInsertId,RowsAffected)
	return rsp, nil
}

//select
func AcessMysqlSelect(ctx context.Context) (rsp string,err error){
	var users []model.User
	// 读取所有字段，select字段尽量只select自己关心的字段，不要用*
	sqlStr := fmt.Sprintf("select * from %s WHERE name = ? OR name = ?\"", defaultTableName)
	err = mysqlClientProxy.QueryToStructs(ctx, &users, sqlStr, "小明", "小红")
	if err != nil {
		err = fmt.Errorf("select error, err:%+v", err)
		return "",err
	}
	for i,user := range users{
		rsp += fmt.Sprintf("%d : %+v",i,user)
	}
	return rsp, nil
}

//delete
func AcessMysqlDelete(ctx context.Context) (rsp string,err error){
	sqlStr := fmt.Sprintf("delete from %s WHERE name = ?", defaultTableName)
	result, err := mysqlClientProxy.Exec(ctx, sqlStr,"小红")
	if err != nil {
		err = fmt.Errorf("delete error, err:%+v", err)
		return "",err
	}
	LastInsertId,_ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected,_ :=result.RowsAffected()  // 被影响的行数
	rsp = fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ",LastInsertId,RowsAffected)
	return rsp, nil
}

//update
func AcessMysqlUpdate(ctx context.Context) (rsp string,err error){
	sqlStr := fmt.Sprintf("update %s set age = ? WHERE id = ?", defaultTableName)
	result, err := mysqlClientProxy.Exec(ctx, sqlStr, 20, 1)
	if err != nil {
		err = fmt.Errorf("update error, err:%+v", err)
		return "",err
	}
	LastInsertId,_ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected,_ :=result.RowsAffected()  // 被影响的行数
	rsp = fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ",LastInsertId,RowsAffected)
	return rsp, nil
}