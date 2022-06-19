package userService

import "ransmart_auth/app/models/userModel"

type IUserService interface {
	FindAll() ([]userModel.User, error)
	FindByID(id int) (userModel.User, error)
	FindByUsername(username string) (userModel.User, error)
	Create(user userModel.User) (err error)
	Update(id int, User userModel.User) (userModel.User, error)
	UpdateV2(user userModel.User) (userModel.User, error)
	Delete(userModel.User) (userModel.User, error)
}
