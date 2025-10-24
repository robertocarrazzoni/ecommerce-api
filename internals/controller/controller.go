package controller

import (
	"github.com/robertocarrazzoni/ecommerce-api/internals/repository"
)

type Controller struct {
	repo *repository.Repository
}

func NewController(repo *repository.Repository) *Controller {
	return &Controller{repo: repo}
}