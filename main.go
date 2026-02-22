package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ----- 需要修改的部分：改成你自己的信息 -----
	// 格式：用户名:密码@tcp(公网地址:端口)/数据库名?参数
	dsn := "testuser:SXX@01743t@tcp(rm-uf67513l4k07d60ar1o.mysql.rds.aliyuncs.com:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	// ---------------------------------------

	// 尝试连接数据库
	fmt.Println("正在尝试连接阿里云数据库...")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接参数错误: ", err)
	}
	defer db.Close()

	// Ping一下，真正建立连接
	err = db.Ping()
	if err != nil {
		log.Fatal("连接失败，请检查网络或账号密码: ", err)
	}

	fmt.Println("恭喜！Go 代码成功连上了阿里云数据库！")

	// 可选：查一下数据库版本确认一下
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal("查询失败: ", err)
	}
	fmt.Println("数据库版本是: ", version)
}
// 同事添加的注释
