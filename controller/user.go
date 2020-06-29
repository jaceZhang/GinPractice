package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gotest/common"
	"gotest/model"
	"gotest/mysql"
)
//用户注册
func Register(c *gin.Context)  {
	var user model.UserModel
	if err:=c.ShouldBind(&user); err!=nil{
		common.OutJson(c,0,user.GetError(err.(validator.ValidationErrors))[0],[]int{})
		return
	}

	mysql.Db.Create(&user)

	result:=mysql.Db.NewRecord(user)

	if result == false{
		common.OutJson(c,1,"注册成功!",[]int{})
	}else{
		common.OutJson(c,0,"注册失败!",[]int{})
	}
}
//获取记录数量
func GetCount(c *gin.Context){
	var count int;
	mysql.Db.Table("user").Count(&count)
	common.OutJson(c,1,"获取成功!", map[string]interface{}{"num":count})
}
//登录
func Login(c *gin.Context)  {
	var user model.UserModel
	var check model.UserModel
	if err:=c.ShouldBind(&user); err!=nil{
		common.OutJson(c,0,user.GetError(err.(validator.ValidationErrors))[0],[]int{})
		return
	}
	mysql.Db.Where(&user).Select("email,password").Find(&check)

	if (check.Email == user.Email) && (check.Password == user.Password){
		fmt.Printf("%v",user)
		fmt.Printf("%v",check)
		common.OutJson(c,1,"登录成功!",[]int{})
	}else{
		common.OutJson(c,0,"账号或密码错误!",[]int{})
	}

}

//删除用户
func DelUser(c *gin.Context){
	id := c.Query("id")
	if id == ""{
		common.OutJson(c,0,"请传入查询ID!",[]int{})
		return
	}
	err := mysql.Db.Where("id = ?",id).Unscoped().Delete(model.UserModel{}).Error
	if err != nil{
		common.OutJson(c,0,"删除失败!",[]int{})
	}else{
		common.OutJson(c,1,"删除成功!",[]int{})
	}
}
