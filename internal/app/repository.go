package app

import (
	"github.com/AliRasoulinejad/hetab/internal/repository"
)

type Repositories struct {
	OrderRepo repository.Order
}

func (application *Application) WithRepositories() {
	application.Repositories = new(Repositories)
	application.Repositories.OrderRepo = repository.NewOrder(application.DB)
}
