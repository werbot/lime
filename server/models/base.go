package models

import (
	"errors"

	"github.com/werbot/lime/config"
)

var (
	ErrKeyNotFound      = errors.New("Key Not Found")
	ErrLicenseNotFound  = errors.New("License Not Found")
	ErrTarifNotFound    = errors.New("Tarif Not Found")
	ErrCustomerNotFound = errors.New("Customer Not Found")
)

func init() {
	config.DB.Debug().AutoMigrate(&Tarif{}, &Customer{}, &Subscription{}, &License{})
}
