package main

import (
	"fmt"
	//前面加_表示只执行init方法
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	//go get "github.com/go-sql-driver/mysql"
	//go get "github.com/jmoiron/sqlx"

	//打开数据库 db, err := sqlx.Open("mysql","username:password@tcp(ip:port)/database?charset=utf8")
	db, err := connectMysql()
	if err != nil {
		fmt.Println(err)
		return
	}
	//db.SetMaxIdleConns(0) //<=0 is mean no limit
	//db.SetMaxOpenConns(0)
	//db.SetConnMaxLifetime(0)
	defer db.Close()
	//exec函数用来增删改
	result, err := db.Exec("update User set name='dwx1' where id = 1")
	result, err = db.Exec("insert into user  values(?,?)", 0, "johny")
	//delete from User where id = 2
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(affected)
	//query函数用来查询
	rows, err := db.Query("select * from User where id <10")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println(id, name)
	}
	//关闭结果集（释放连接）
	rows.Close()
}

var (
	userName  = "root"
	password  = "123456"
	ipAddrees = "localhost"
	port      = 3306
	dbName    = "testdb"
	charset   = "utf8"
)

func connectMysql() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	return sqlx.Open("mysql", dsn) //执行db的方法时候才会真正从连接池中建立获取连接
}
