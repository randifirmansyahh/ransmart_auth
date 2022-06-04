package userModel

import "ransmart_auth/app/helper/modelHelper"

type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Firstname string `json:"firstname" validate:"required,min=3,max=50,alpha"`
	Lastname  string `json:"lastname" validate:"required,min=3,max=50,alpha"`
	Username  string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password  string `json:"password" validate:"required,min=3,max=50"`
	Email     string `json:"email" validate:"required,email,min=3,max=50"`
	No_Hp     string `json:"no_hp" validate:"required,min=3,max=50,numeric"`
	Image     string `json:"image" validate:"required,min=3,max=1000,url"`
	modelHelper.DateAuditModel
}

type ReqUserLogin struct {
	Username string `json:"username" validate:"required,min=3,max=50,alphanum"`
	Password string `json:"password" validate:"required,min=3,max=50"`
}
