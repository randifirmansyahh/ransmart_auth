package userRepository

import (
	"ransmart_auth/app/models/userModel"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() ([]userModel.User, error)
	FindByID(id int) (userModel.User, error)
	FindByUsername(username string) (userModel.User, error)
	Create(tx *gorm.DB, user userModel.User) (err error)
	UpdateV2(user userModel.User) (userModel.User, error)
	Update(id int, user userModel.User) (userModel.User, error)
	Delete(user userModel.User) (userModel.User, error)
}
