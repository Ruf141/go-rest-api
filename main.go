package main

import (
	"go-rest-api/db"
	"go-rest-api/repository"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
}
