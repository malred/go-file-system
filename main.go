package main

import ( 
	db "go_file_db"
	web "go_file_web"
)

func main() {
	// 初始化数据库
	db.InitDB()   
	// 开启服务
	web.Run()
}
