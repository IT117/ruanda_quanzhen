package main

import
(
	"DetaCerProject/blockchain"
	"DetaCerProject/db_mysql"
	_ "DetaCerProject/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1.生成第一个区块链
	block :=blockchain.CreateGenesisBlock()
	fmt.Println("第一个区块：",block)
	fmt.Printf("区块的Hash值：%x",block.Hash)
	return
	//连接数据库
	db_mysql.ConnectDB()
	//静态资源路径设置
	beego.SetStaticPath("js","./static/js")
	beego.SetStaticPath("css","./static/css")
	beego.SetStaticPath("img","./static/img")

	beego.Run()
}

