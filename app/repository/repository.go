package repository

import (
	"ransmart_auth/app/repository/userRepository"
)

type Repository struct {
	IUserRepository userRepository.IUserRepository
}
