package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	Email string `form:"email" binding:"required,alphanum,min=4"`
	Password string `form:"password" binding:"required,min=6,max=12"`
	gorm.Model
}

func (UserModel) TableName() string  {
	return "user"
}
//自定义参数验证
func (u *UserModel) GetError (err validator.ValidationErrors) (msg []string){
	for _,v:=range err{
		if v.Field() == "Email"{
			switch v.Tag() {
				case "required":
					msg=append(msg,"邮箱号不能为空!")
				case "min":
					msg=append(msg,"邮箱号最小长度为4!")
				case "alphanum":
					msg=append(msg,"邮箱号只能包含字母和数字!")
			}
		}
		if v.Field() == "Password"{
			switch v.Tag() {
				case "required":
					msg=append(msg,"密码不能为空!")
				case "min":
					msg=append(msg,"密码最小长度为6位!")
				case "max":
					msg=append(msg,"密码最大长度为12位!")
			}
		}
	}
	return
}

