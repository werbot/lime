package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/config"
)

type Customer struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Customer) SaveCustomer() (*Customer, error) {
	err := config.DB.Debug().Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

func (c *Customer) FindCustomerByID(uid uint32) (*Customer, error) {
	err := config.DB.Debug().Model(Customer{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, ErrCustomerNotFound
	}
	return c, err
}

func (c *Customer) UpdateCustomer(uid uint32) (*Customer, error) {
	db := config.DB.Debug().Model(&Customer{}).Where("id = ?", uid).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"name":      c.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Customer{}, db.Error
	}

	err := db.Debug().Model(&Customer{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

func (c *Customer) DeleteCustomer(uid uint32) (int64, error) {
	db := config.DB.Debug().Model(&Customer{}).Where("id = ?", uid).Take(&Customer{}).Delete(&Customer{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
