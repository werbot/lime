package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/config"
)

// Customer is a ...
type Customer struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Status    bool      `gorm:"false" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveCustomer is a ...
func (c *Customer) SaveCustomer() (*Customer, error) {
	err := config.DB.Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

// FindCustomerByID is a ...
func (c *Customer) FindCustomerByID(uid uint32) (*Customer, error) {
	err := config.DB.Model(Customer{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, ErrCustomerNotFound
	}
	return c, err
}

// UpdateCustomer is a ...
func (c *Customer) UpdateCustomer(uid uint32) (*Customer, error) {
	db := config.DB.Model(&Customer{}).Where("id = ?", uid).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"name":      c.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Customer{}, db.Error
	}

	err := db.Model(&Customer{}).Where("id = ?", uid).Take(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}

// DeleteCustomer is a ...
func (c *Customer) DeleteCustomer(uid uint32) (int64, error) {
	db := config.DB.Model(&Customer{}).Where("id = ?", uid).Take(&Customer{}).Delete(&Customer{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// CustomersList is a ...
func CustomersList() *[]Customer {
	customers := []Customer{}
	db := config.DB.Find(&customers)
	if db.Error != nil {
		return &customers
	}
	return &customers
}
