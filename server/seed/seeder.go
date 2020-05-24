package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/server/models"
)

var tarifs = []models.Tarif{
	models.Tarif{
		Name:      "Garage",
		Price:     590,
		Servers:   10,
		Companies: 10,
		Users:     10,
	},
	models.Tarif{
		Name:      "Startup",
		Price:     1990,
		Servers:   20,
		Companies: 20,
		Users:     20,
	},
}

var customers = []models.Customer{
	models.Customer{
		ID:   1,
		Name: "Alex Past",
	},
	models.Customer{
		ID:   2,
		Name: "Lisa Boston",
	},
	models.Customer{
		ID:   3,
		Name: "Adam Potar",
	},
	models.Customer{
		ID:   4,
		Name: "Greg Gordon",
	},
}

var subscription = []models.Subscription{
	models.Subscription{
		CustomerID: 1,
		StripeID:   "cus_FEDaLVeqQoVy6m",
		TarifID:    1,
		Status:     true,
	},
	models.Subscription{
		CustomerID: 2,
		StripeID:   "cus_APBaLDeqQoVy8m",
		TarifID:    2,
		Status:     true,
	},
	models.Subscription{
		CustomerID: 3,
		StripeID:   "cus_FEDajfeqSkTy01",
		TarifID:    2,
		Status:     true,
	},
	models.Subscription{
		CustomerID: 4,
		StripeID:   "cus_GGjdLDeqQokfj5",
		TarifID:    1,
		Status:     false,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Tarif{}, &models.Customer{}, &models.Subscription{}, &models.License{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Tarif{}, &models.Customer{}, &models.Subscription{}, &models.License{}).Error

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range tarifs {
		err = db.Debug().Model(&models.Tarif{}).Create(&tarifs[i]).Error
		if err != nil {
			log.Fatalf("cannot seed tarif table: %v", err)
		}
	}

	for i, _ := range customers {
		err = db.Debug().Model(&models.Customer{}).Create(&customers[i]).Error
		if err != nil {
			log.Fatalf("cannot seed customer table: %v", err)
		}
	}

	for i, _ := range subscription {
		err = db.Debug().Model(&models.Subscription{}).Create(&subscription[i]).Error
		if err != nil {
			log.Fatalf("cannot seed subscription table: %v", err)
		}
	}
}
