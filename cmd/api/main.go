package main

import (
	"github.com/gin-gonic/gin"
	db "github.com/robertocarrazzoni/ecommerce-api/internals"
	"github.com/robertocarrazzoni/ecommerce-api/internals/controller"
	"github.com/robertocarrazzoni/ecommerce-api/internals/repository"
)

func main() {
	r := gin.Default()

	db.Connect()

	repo := repository.NewRepository(db.DB)
	controller.NewController(repo)

	r.Run()

}
