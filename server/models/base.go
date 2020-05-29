package models

import (
	"errors"

	"github.com/werbot/lime/config"
)

var (
	// ErrKeyNotFound is a ...
	ErrKeyNotFound = errors.New("Key Not Found")

	// ErrLicenseNotFound is a ...
	ErrLicenseNotFound = errors.New("License Not Found")

	// ErrTariffNotFound is a ...
	ErrTariffNotFound = errors.New("Tariff Not Found")

	// ErrCustomerNotFound is a ...
	ErrCustomerNotFound = errors.New("Customer Not Found")
)

func init() {
	config.DB.AutoMigrate(&Tariff{}, &Customer{}, &Subscription{}, &License{})
}
