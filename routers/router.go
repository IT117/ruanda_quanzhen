package routers

import (
"DetaCerProject/controllers"
"github.com/astaxie/beego"
)
/**
rouer.go文件的作用：路由功能。用于接收并分发接收到的浏览器的请求，并分配请求
*/

func init() {
	beego.Router("/", &controllers.MainController{})
	//用户注册接口请求
	beego.Router("/user_register",&controllers.RegisterController{})
	//直接登录页面的接口请求
	beego.Router("/login.html",&controllers.LoginController{})
	//用户登录接口
	beego.Router("/user_login",&controllers.LoginController{})
	//用户传证接口
	beego.Router("/upload",&controllers.UploadController{})
	//立即存在接口
	beego.Router("/list_record",&controllers.UploadController{})

}
