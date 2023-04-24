package go_file_db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// 数据库相关操作
var db *sqlx.DB

// 初始化数据库连接
func InitDB() (err error) {
	dsn := "./fileSystem.db"
	// 连接
	// Open可能仅校验参数，而没有与db间创建连接，
	// 要确认db是否可用，需要调用Ping。Connect则相当于Open+Ping。
	db, err = sqlx.Connect("sqlite3", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	// 最大连接数
	db.SetMaxOpenConns(100)
	// 最大空闲连接数
	db.SetMaxIdleConns(16) 
	return
}
 