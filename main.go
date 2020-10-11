package main

import
(
	"DetaCerProject/db_mysql"
	_ "DetaCerProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//连接数据库
	db_mysql.ConnectDB()
	//静态资源路径设置
	beego.SetStaticPath("js","./static/js")
	beego.SetStaticPath("css","./static/css")
	beego.Run()
}

