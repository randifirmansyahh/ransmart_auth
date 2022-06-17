package loginHandler

import (
	"encoding/json"
	"log"
	"net/http"
	"ransmart_auth/app/helper/helper"
	"ransmart_auth/app/helper/response"
	"ransmart_auth/app/helper/tokenHelper"
	"ransmart_auth/app/models/tokenModel"
	"ransmart_auth/app/models/userModel"
	"ransmart_auth/app/service"
	"strconv"

	"github.com/go-playground/validator"
)

var (
	WAKTU        = tokenHelper.WAKTU
	AUD          = tokenHelper.AUD
	ISS          = tokenHelper.ISS
	LOGIN_SECRET = tokenHelper.LOGIN_SECRET
)

type loginHandler struct {
	service service.Service
}

func NewLoginHandler(loginService service.Service) *loginHandler {
	return &loginHandler{loginService}
}

func (l *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	// cek user dan pass
	// decode from json
	decoder := json.NewDecoder(r.Body)

	// fill to model
	var datarequest userModel.ReqUserLogin
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Gagal login", nil)
		return
	}

	validate := validator.New()
	err := validate.Struct(datarequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			response.Response(w, http.StatusBadRequest, errors.Error(), nil)
			return
		}
	}

	// select ke db
	cari, err := l.service.IUserService.FindByUsername(datarequest.Username)
	if err != nil {
		response.Response(w, http.StatusOK, "Username tidak ditemukan", nil)
		return
	}

	// hash password from request
	reqPwd := helper.Encode([]byte(datarequest.Password))
	CariPwd, err := helper.Decode([]byte(cari.Password))
	if err != nil {
		response.Response(w, http.StatusBadRequest, "Gagal login", nil)
		return
	}

	log.Println(string(CariPwd))
	log.Println(string(reqPwd))

	// bandingkan
	if cari.Username != datarequest.Username || string(CariPwd) != string(reqPwd) {
		response.Response(w, http.StatusOK, "Password salah", nil)
		return
	}

	// buat expired time nya
	// convert
	newWaktu, _ := strconv.Atoi(WAKTU)
	expiredTime := helper.ExpiredTime(newWaktu)

	// fill ke jwt
	token, err := tokenHelper.BuatJWT(ISS, AUD, LOGIN_SECRET, expiredTime)

	// cek jika generate gagal
	if err != nil {
		response.Response(w, http.StatusInternalServerError, "Gagal login", nil)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token

	response.Response(w, http.StatusOK, "Login berhasil", tokensMaps)
}

func (l *loginHandler) Register(w http.ResponseWriter, r *http.Request) {
	// cek user dan pass
	// decode from json
	decoder := json.NewDecoder(r.Body)

	// fill to model
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		response.Response(w, http.StatusBadRequest, "Gagal register", nil)
		return
	}

	validate := validator.New()
	err := validate.Struct(datarequest)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if errors != nil {
			response.Response(w, http.StatusBadRequest, errors.Error(), nil)
			return
		}
	}

	// select ke db
	cari, err := l.service.IUserService.FindByUsername(datarequest.Username)
	if cari.Username == datarequest.Username && err == nil {
		response.Response(w, http.StatusOK, "Username tersebut telah digunakan", nil)
		return
	}

	// hash password from request
	newPassword := helper.Encode([]byte(datarequest.Password))
	datarequest.Password = string(newPassword)

	// insert ke db
	_, err = l.service.IUserService.Create(datarequest)
	if err != nil {
		response.Response(w, http.StatusInternalServerError, "Gagal register", nil)
		return
	}

	response.Response(w, http.StatusOK, "Register berhasil", nil)
}

func (l *loginHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	// convert
	newWaktu, _ := strconv.Atoi(WAKTU)

	// buat expired time nya
	expiredTime := helper.ExpiredTime(newWaktu)

	// fill ke jwt
	token, err := tokenHelper.BuatJWT(ISS, AUD, LOGIN_SECRET, expiredTime)

	// cek jika generate gagal
	if err != nil {
		response.Response(w, http.StatusInternalServerError, "Gagal login", nil)
		return
	}

	// masukin ke model, kirim respon
	var tokensMaps tokenModel.Token
	tokensMaps.FullToken = token
	response.Response(w, http.StatusOK, "Login berhasil", tokensMaps)
}
