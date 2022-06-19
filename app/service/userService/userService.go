package userService

import (
	"encoding/json"
	"errors"
	"os"
	"ransmart_auth/app/helper/helper"
	"ransmart_auth/app/helper/httpRequest"
	"ransmart_auth/app/helper/tokenHelper"
	"ransmart_auth/app/models/userModel"
	"ransmart_auth/app/repository"
	"strconv"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type service struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewService(repository repository.Repository, db *gorm.DB) *service {
	return &service{repository, db}
}

func (s *service) FindAll() ([]userModel.User, error) {
	return s.repository.IUserRepository.FindAll()
}

func (s *service) FindByID(id int) (userModel.User, error) {
	return s.repository.IUserRepository.FindByID(id)
}

func (s *service) FindByUsername(username string) (userModel.User, error) {
	return s.repository.IUserRepository.FindByUsername(username)
}

func (s *service) Create(user userModel.User) (err error) {
	// hashing password
	newPassword := helper.Encode([]byte(user.Password))
	user.Password = string(newPassword)

	// create user to database
	tx := s.db.Begin()
	err = s.repository.IUserRepository.Create(tx, user)
	if err != nil {
		return errors.New("gagal menambahkan user")
	}

	// Get environment variable
	ISS := os.Getenv("JWT_ISS")
	AUD := os.Getenv("JWT_AUD")
	JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")
	JWT_EXPIRATION_DURATION_DAY := os.Getenv("JWT_EXPIRATION_DURATION_DAY")
	newWaktu, _ := strconv.Atoi(JWT_EXPIRATION_DURATION_DAY)
	expiredTime := helper.ExpiredTime(newWaktu)

	// create token
	jwt, err := tokenHelper.BuatJWT(ISS, AUD, JWT_SECRET_KEY, expiredTime)
	if err != nil {
		log.Error().Msgf("error buat jwt: %v", err)
		return errors.New("gagal generate token")
	}

	// set header
	header := map[string]string{
		"Authorization": "Bearer " + jwt,
	}

	// User Payload
	userPayload := userModel.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		No_Hp:     user.No_Hp,
		Image:     user.Image,
	}

	// set payload to byte
	userPayloadByte, _ := json.Marshal(userPayload)

	// add waitgroup and error count
	var (
		wg       sync.WaitGroup
		errCount int
	)

	// ransmart_product
	wg.Add(1)
	go func() {
		defer wg.Done()
		urlProduct := "https://ransmart-product.herokuapp.com/user"
		code, _, err := httpRequest.HTTPResponse("POST", urlProduct, string(userPayloadByte), header)
		if err != nil || code != 200 {
			tx.Rollback()
			log.Error().Msgf("error create user to ransmart_product : %v", err)
			errCount++
		}
	}()

	// ransmart_pay
	wg.Add(1)
	go func() {
		defer wg.Done()
		urlProduct := "https://ransmart-pay.herokuapp.com/user"
		code, _, err := httpRequest.HTTPResponse("POST", urlProduct, string(userPayloadByte), header)
		if err != nil || code != 200 {
			tx.Rollback()
			log.Error().Msgf("error create user to ransmart_pay : %v", err)
			errCount++
		}
	}()

	// wait for all goroutine
	wg.Wait()

	// check error
	if errCount > 0 {
		tx.Rollback()
		log.Error().Msgf("error create user to ransmart_product and ransmart_pay : %v", err)
		return errors.New("gagal menambahkan user")
	}

	// commit transaction
	err = tx.Commit().Error
	log.Info().Msgf("error commit user : %v", err)

	return
}

func (s *service) Update(id int, User userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.Update(id, User)
}

func (s *service) UpdateV2(user userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.UpdateV2(user)
}

func (s *service) Delete(data userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.Delete(data)
}
