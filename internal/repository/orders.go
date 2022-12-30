package repository

import (
	"gorm.io/gorm"
)

type Order interface {
}

type order struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) Order {
	return order{db: db}
}
