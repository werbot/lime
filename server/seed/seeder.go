package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/server/models"
)

var tariffs = []models.Tariff{
	{
		Name:      "Garage",
		Price:     590,
		Servers:   10,
		Companies: 10,
		Users:     10,
	},
	{
		Name:      "Startup",
		Price:     1990,
		Servers:   20,
		Companies: 20,
		Users:     20,
	},
}

var customers = []models.Customer{
	{
		ID:     1,
		Name:   "Alex Past",
		Status: true,
	},
	{
		ID:     2,
		Name:   "Lisa Boston",
		Status: true,
	},
	{
		ID:     3,
		Name:   "Adam Potar",
		Status: true,
	},
	{
		ID:     4,
		Name:   "Greg Gordon",
		Status: true,
	},
}

var subscription = []models.Subscription{
	{
		CustomerID: 1,
		StripeID:   "cus_FEDaLVeqQoVy6m",
		TariffID:   1,
		Status:     true,
	},
	{
		CustomerID: 2,
		StripeID:   "cus_APBaLDeqQoVy8m",
		TariffID:   2,
		Status:     true,
	},
	{
		CustomerID: 3,
		StripeID:   "cus_FEDajfeqSkTy01",
		TariffID:   2,
		Status:     true,
	},
	{
		CustomerID: 4,
		StripeID:   "cus_GGjdLDeqQokfj5",
		TariffID:   1,
		Status:     false,
	},
}

// Load import test data to database
func Load(db *gorm.DB) {
	err := db.DropTableIfExists(&models.Tariff{}, &models.Customer{}, &models.Subscription{}, &models.License{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.AutoMigrate(&models.Tariff{}, &models.Customer{}, &models.Subscription{}, &models.License{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range tariffs {
		err = db.Model(&models.Tariff{}).Create(&tariffs[i]).Error
		if err != nil {
			log.Fatalf("cannot seed tariff table: %v", err)
		}
	}

	for i := range customers {
		err = db.Model(&models.Customer{}).Create(&customers[i]).Error
		if err != nil {
			log.Fatalf("cannot seed customer table: %v", err)
		}
	}

	for i := range subscription {
		err = db.Model(&models.Subscription{}).Create(&subscription[i]).Error
		if err != nil {
			log.Fatalf("cannot seed subscription table: %v", err)
		}
	}
}
