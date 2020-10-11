package models

import (
	"DetaCerProject/db_mysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int `form:"id"`
	Phone string  `form:"phone"`
	Password string `form:"password"`

}
//保存数据到数据库的方法
func (u User) SaveUser()(int64,error){
	//1.密码脱敏处理
	hashMd5:=md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes:=hashMd5.Sum(nil)
	u.Password =hex.EncodeToString(bytes)

	//保存数据到数据库
	row ,err:=db_mysql.DB.Exec("insert into user (phone,password) " +"value (?,?)",u.Phone,u.Password)
	if err!=nil {
		return -1,err

	}
	id,err:=row.RowsAffected()
	if err!=nil {
		return -1,err}
	return id,nil

}
func (u User)QueryUser() (*User, error) {
	row:=db_mysql.DB.QueryRow("select phone from user where  phone=? and  password=?",u.Phone,u.Password)

    err:=row.Scan(&u.Phone)
	if err!=nil {
		return nil ,err


	}
	return &u,nil
}



