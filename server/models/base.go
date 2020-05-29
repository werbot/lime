package models

import (
	"errors"

	"github.com/werbot/lime/config"
)

var (
	ErrKeyNotFound      = errors.New("Key Not Found")
	ErrLicenseNotFound  = errors.New("License Not Found")
	ErrTariffNotFound   = errors.New("Tariff Not Found")
	ErrCustomerNotFound = errors.New("Customer Not Found")
)

func init() {
	config.DB.AutoMigrate(&Tariff{}, &Customer{}, &Subscription{}, &License{})
}
