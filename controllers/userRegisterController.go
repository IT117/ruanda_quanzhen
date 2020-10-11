package controllers

import (
	"DetaCerProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r*RegisterController) Post() {
	//1.解析请求数据
	var user  models.User
	err :=r.ParseForm(&user)
	if err != nil {
		//2返回用户信息给浏览器,提示用户
		r.Ctx.WriteString("抱歉用户解析数据错误，请重试")
		return

	}
	//保存数据到数据库
	_,err=user.SaveUser()
	if err!=nil {


		r.Ctx.WriteString("抱歉用户注册失败，请重试")
		return

	}
	r.TplName="login.html"





}