package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"git.code.oa.com/trpc-go/trpc-database/mysql"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/log"

	"git.code.oa.com/video_app_short_video/hello_alice/common"
	"git.code.oa.com/video_app_short_video/hello_alice/model"
	"git.code.oa.com/video_app_short_video/hello_alice/config"
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
func InitMysqlProxy() error{
	// 获取配置信息
	serviceConfig := config.GetConfig()
	mysqlConfig := serviceConfig.Mysql
	target := fmt.Sprintf("dsn://%s:%s@tcp(%s:%d)/%s",
		mysqlConfig.User,mysqlConfig.Password,mysqlConfig.Domain,mysqlConfig.Port,mysqlConfig.DB)
	mysqlClientProxy = mysql.NewClientProxy(
		mysqlConfig.ServiceName,
		//dsn://user:passwd@tcp(vip:port)/db?timeout=1s&parseTime=true&interpolateParams=true")  mdb使用域名多实例需要加上 &interpolateParams=true
		client.WithTarget(target +"?timeout=1s&parseTime=true&interpolateParams=true"),
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
//create 创建表
func AcessMysqlInit(ctx context.Context) (rsp string,err error) {
	sqlStr := fmt.Sprintf("create table if not exists %s (`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(30),`age` int(11),PRIMARY KEY (`id`));", defaultTableName)
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
		result := common.AttaSendFields(fmt.Sprintf("%v",err), "AcessMysqlInsert error")
		log.Infof("AcessMysqlInsert atta SendString(), result:" + strconv.Itoa(result))
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
	users = make([]model.User, 0)
	// 框架底层自动for循环rows调用该next函数
	var user model.User
	next := func(rows *sql.Rows) error {
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			return err
		}
		users = append(users, user)
		return nil
	}
	// 读取所有字段，select字段尽量只select自己关心的字段，不要用*
	sqlStr = fmt.Sprintf("select * from %s limit ?", defaultTableName)
	err = mysqlClientProxy.Query(ctx, next, sqlStr, 100)
	if err != nil {
		err = fmt.Errorf("select error, err:%+v", err)
		return "",err
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