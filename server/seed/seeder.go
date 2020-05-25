package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/server/models"
)

var tarifs = []models.Tarif{
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
		ID:   1,
		Name: "Alex Past",
	},
	{
		ID:   2,
		Name: "Lisa Boston",
	},
	{
		ID:   3,
		Name: "Adam Potar",
	},
	{
		ID:   4,
		Name: "Greg Gordon",
	},
}

var subscription = []models.Subscription{
	{
		CustomerID: 1,
		StripeID:   "cus_FEDaLVeqQoVy6m",
		TarifID:    1,
		Status:     true,
	},
	{
		CustomerID: 2,
		StripeID:   "cus_APBaLDeqQoVy8m",
		TarifID:    2,
		Status:     true,
	},
	{
		CustomerID: 3,
		StripeID:   "cus_FEDajfeqSkTy01",
		TarifID:    2,
		Status:     true,
	},
	{
		CustomerID: 4,
		StripeID:   "cus_GGjdLDeqQokfj5",
		TarifID:    1,
		Status:     false,
	},
}

// Load import test data to database
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Tarif{}, &models.Customer{}, &models.Subscription{}, &models.License{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Tarif{}, &models.Customer{}, &models.Subscription{}, &models.License{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range tarifs {
		err = db.Debug().Model(&models.Tarif{}).Create(&tarifs[i]).Error
		if err != nil {
			log.Fatalf("cannot seed tarif table: %v", err)
		}
	}

	for i := range customers {
		err = db.Debug().Model(&models.Customer{}).Create(&customers[i]).Error
		if err != nil {
			log.Fatalf("cannot seed customer table: %v", err)
		}
	}

	for i := range subscription {
		err = db.Debug().Model(&models.Subscription{}).Create(&subscription[i]).Error
		if err != nil {
			log.Fatalf("cannot seed subscription table: %v", err)
		}
	}
}
