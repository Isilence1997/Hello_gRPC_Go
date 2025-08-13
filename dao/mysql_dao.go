package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"git.code.oa.com/gRPC-go/gRPC-database/mysql"
	"git.code.oa.com/gRPC-go/gRPC-go/client"
	"git.code.oa.com/gRPC-go/gRPC-go/log"

	"git.code.oa.com/video_app_short_video/hello_alice/common"
	"git.code.oa.com/video_app_short_video/hello_alice/config"
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

// 初始化mysql
func InitMysqlProxy() error {
	// 获取配置信息
	serviceConfig := config.GetConfig()
	mysqlConfig := serviceConfig.Mysql
	target := fmt.Sprintf("dsn://%s:%s@tcp(%s:%d)/%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Domain, mysqlConfig.Port, mysqlConfig.DB)
	mysqlClientProxy = mysql.NewClientProxy(
		mysqlConfig.ServiceName,
		//dsn://user:passwd@tcp(vip:port)/db?timeout=1s&parseTime=true&interpolateParams=true")  mdb使用域名多实例需要加上 &interpolateParams=true
		client.WithTarget(target+"?timeout=1s&parseTime=true&interpolateParams=true"),
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
	err := mysqlClientProxy.Query(context.Background(), next, "SELECT 1") //增加临时列，列值为1
	return err
}

/*
//create 创建表
func AcessMysqlInit(ctx context.Context) (rsp string,err error) {
	sqlStr := fmt.Sprintf("create table if not exists %s (`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(30),`age` int(11),PRIMARY KEY (`id`));", defaultTableName)
	//调用 mysqlClientProxy.Exec 执行建表语句
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
*/
// insert
func AcessMysqlInsert(ctx context.Context) (rsp string, err error) {
	//组装sql：往 defaultTableName 表一次插入两行（(?,?),(?,?)）。
	sqlStr := fmt.Sprintf("insert into %s(name,age) values (?,?),(?,?)", defaultTableName)
	//用参数绑定传入两行 name/age（小明，17）（小红，18）
	result, err := mysqlClientProxy.Exec(ctx, sqlStr, "小明", 17, "小红", 18)
	if err != nil {
		err = fmt.Errorf("insert error, err:%+v", err)
		attaCode := common.AttaSendFields(fmt.Sprintf("%v", err), "AcessMysqlInsert error")
		log.Infof("AcessMysqlInsert atta SendString(), result:" + strconv.Itoa(attaCode))
		return "", err
	}
	LastInsertId, _ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected, _ := result.RowsAffected() // 被影响的行数
	rsp = fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ", LastInsertId, RowsAffected)
	return rsp, nil
}

// select
func AcessMysqlSelect(ctx context.Context) (rsp string, err error) {
	//method1
	var users []model.User
	// 读取所有字段，select字段尽量只select自己关心的字段，不要用*
	sqlStr := fmt.Sprintf("select * from %s WHERE name = ? OR name = ?", defaultTableName)
	//QueryToStructs自动执行 SQL，并将结果按字段映射填充到 users中。
	err = mysqlClientProxy.QueryToStructs(ctx, &users, sqlStr, "小明", "小红")
	if err != nil {
		err = fmt.Errorf("select error, err:%+v", err)
		return "", err
	}
	for i, user := range users {
		log.Infof("%d : %+v", i, user)
		rsp = "OK"
	}
	//method2
	users = make([]model.User, 0)
	//定义next回调函数
	next := func(rows *sql.Rows) error {
		var user model.User
		//把当前行的数据映射到 user 的字段。
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return err
		}
		users = append(users, user)
		return nil
	}
	sqlStr = fmt.Sprintf("select * from %s limit ?", defaultTableName)
	//Query 会执行 SQL 并在内部循环调用 next，直到数据读完。
	err = mysqlClientProxy.Query(ctx, next, sqlStr, 100)
	if err != nil {
		err = fmt.Errorf("select error, err:%+v", err)
		return "", err
	}
	return rsp, nil
}

// delete
func AcessMysqlDelete(ctx context.Context) (rsp string, err error) {
	sqlStr := fmt.Sprintf("delete from %s WHERE name = ?", defaultTableName)
	result, err := mysqlClientProxy.Exec(ctx, sqlStr, "小红")
	if err != nil {
		err = fmt.Errorf("delete error, err:%+v", err)
		return "", err
	}
	LastInsertId, _ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected, _ := result.RowsAffected() // 被影响的行数
	rsp = fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ", LastInsertId, RowsAffected)
	return rsp, nil
}

// update
func AcessMysqlUpdate(ctx context.Context) (rsp string, err error) {
	sqlStr := fmt.Sprintf("update %s set age = ? WHERE id = ?", defaultTableName)
	result, err := mysqlClientProxy.Exec(ctx, sqlStr, 20, 1)
	if err != nil {
		err = fmt.Errorf("update error, err:%+v", err)
		return "", err
	}
	LastInsertId, _ := result.LastInsertId() // 最后一条更新数据行的主键id
	RowsAffected, _ := result.RowsAffected() // 被影响的行数
	rsp = fmt.Sprintf("LastInsertId = %d,RowsAffected = %d ", LastInsertId, RowsAffected)
	return rsp, nil
}
