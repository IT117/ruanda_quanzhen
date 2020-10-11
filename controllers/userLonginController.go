package controllers

import (
	"DetaCerProject/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	//设置login.html为模板文件
	l.TplName="login.html"

}
func(l *LoginController)Post(){
	var user models.User
	//解析用户信息
	err :=l.ParseForm(&user)
	if err !=nil{
		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试")
		return
	}
	//查询数据库的用户信息
	u,err:=user.QueryUser()
	if err!=nil {
		l.Ctx.WriteString("用户登录失败，请重试")
		return
	}
	//登录成功页面
	l.Data["Phone"]=u.Phone
	l.TplName="home.html"


}