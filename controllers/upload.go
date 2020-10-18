package controllers

import (
	"DetaCerProject/models"
	"DetaCerProject/uilt"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"time"
)

type UploadController struct{
	beego.Controller
}

func (this *UploadController) Get(){
	phone:=this.GetString("phone")
	this.Data["Phone"]=phone
	this.TplName = "home.html"
}

func (this *UploadController) Post(){
	//1.获取客户端上传的文件以及其他的form表单的信息
	//标题
	fileTitle:=this.Ctx.Request.PostFormValue("upload_title")
	phone:=this.Ctx.Request.PostFormValue("phone")

	f, h, _ := this.GetFile("myfile")//获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":true,
		".jpeg":true,
		".png":true,
	}

	if _,ok := AllowExtMap[ext];!ok{
		this.Ctx.WriteString( "后缀名不符合上传要求" )
		return
	}
	//创建目录
	uploadDir := "static/img/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll( uploadDir , 777)
	if err != nil {
		this.Ctx.WriteString( fmt.Sprintf("%v",err) )
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )

	fileName := fmt.Sprintf("%x",hashName) + ext
	//this.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = this.SaveToFile("myfile", fpath)
	if err != nil {
		this.Ctx.WriteString( fmt.Sprintf("%v",err) )
	}
	//计算文件哈希值
	hashFile,err:=os.Open(uploadDir)
	if err!=nil {
		this.Ctx.WriteString("文件计算失败")
		return

	}
	defer hashFile.Close()
	hash,err:=uilt.MD5HashReader(hashFile)
	//3.将文件上传的记录保存到数据库中
	record:=models.UploadRecord{}
	record.Phone=phone
	record.FileName =h.Filename

	record.CertTime=time.Now().Unix()
	record.FileTitle = fileTitle
	record.FileCert=hash
	record.FileSize=h.Size
	_,err =record.SaveRecord()
	if err !=nil{
		fmt.Println(err)
		this.Ctx.WriteString("抱歉，数据认证错误，请重试！")
		return
	}
	//4.从数据库中读取phone用户对应的所有认证数据记录
	records,err:=models.QuerRecordBuPhone(phone)
	//5.根据文件保存的结果，返回相应的提示信息或者跳转页面
	if err !=nil{
		this.Ctx.WriteString("抱歉，获取认证数据失败，请重试！")
		return
	}
	this.Ctx.WriteString("上传成功")
	fmt.Println(records)
	this.Data["Records"] =record
	this.Data["Phone"] =phone
	this.TplName ="list_record.html"





}