package main

import (
	"ransmart_auth/app/helper/helper"
	"ransmart_auth/app/server"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load("params/.env")
	helper.CheckEnv(err)

	// running server
	server.Execute()
}
