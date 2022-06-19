package server

import (
	"log"
	"net/http"
	"os"
	"ransmart_auth/app/handler/loginHandler"
	"ransmart_auth/app/handler/tokenHandler"
	"ransmart_auth/app/handler/userHandler"
	"ransmart_auth/app/helper/helper"
	"ransmart_auth/app/helper/response"
	"ransmart_auth/app/models/userModel"
	"ransmart_auth/app/repository"
	"ransmart_auth/app/repository/userRepository"
	"ransmart_auth/app/service"
	"ransmart_auth/app/service/userService"

	"github.com/go-chi/chi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Execute() {
	// try connect to database
	log.Println("Connecting to Database...")
	db, err := gorm.Open(mysql.Open(getConnectionString()), &gorm.Config{})
	helper.CheckFatal(err)

	// migrate model to database
	db.AutoMigrate(&userModel.User{})
	log.Println("Database Connected")

	// generate repository
	allRepositories := repository.Repository{
		IUserRepository: userRepository.NewRepository(db),
	}

	// try connect to redis
	log.Println("Connecting to Redis in Background...")
	redis := connectToRedis()

	// generate service
	allServices := service.Service{
		IUserService: userService.NewService(allRepositories, db),
	}

	// generate handler
	user := userHandler.NewUserHandler(allServices, redis)
	login := loginHandler.NewLoginHandler(allServices)

	// router
	r := chi.NewRouter()

	// check service
	r.Group(func(g chi.Router) {
		g.Get("/", func(w http.ResponseWriter, r *http.Request) {
			response.ResponseRunningService(w)
		})
	})

	// // global token
	// r.Group(func(g chi.Router) {
	// 	g.Get("/globaltoken", login.GenerateToken)
	// })

	// login
	r.Group(func(l chi.Router) {
		l.Post("/login", login.Login)
		l.Post("/register", login.Register)
	})

	// user
	r.Group(func(u chi.Router) {
		u.Use(tokenHandler.GetToken) // pelindung token
		u.Get("/user", user.GetSemuaUser)
		u.Get("/user/{id}", user.GetUserByID)
		u.Post("/user", user.PostUser)
		u.Put("/user/{id}", user.UpdateUser)
		u.Delete("/user/{id}", user.DeleteUser)
	})

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Println("Service running on " + host + ":" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Println("Error Starting Service")
	}
}
