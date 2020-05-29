package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/config"
)

// Subscription is a ...
type Subscription struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	StripeID   string    `gorm:"size:18;not null;unique" json:"stripe_id"`
	CustomerID uint32    `sql:"type:int REFERENCES customers(id)" json:"customer_id"`
	TariffID   uint32    `sql:"type:int REFERENCES tariffs(id)" json:"tariff_id"`
	Status     bool      `gorm:"false" json:"status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// SaveSubscription is a ...
func (s *Subscription) SaveSubscription() (*Subscription, error) {
	err := config.DB.Create(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	return s, nil
}

// FindSubscriptionByID is a ...
func (s *Subscription) FindSubscriptionByID(uid uint32) (*Subscription, error) {
	err := config.DB.Model(Subscription{}).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Subscription{}, ErrTariffNotFound
	}
	return s, err
}

// FindSubscriptionByStripeID is a ...
func (s *Subscription) FindSubscriptionByStripeID(stripe_id string) (*Subscription, error) {
	err := config.DB.Model(Subscription{}).Where("stripe_id = ? AND status = ?", stripe_id, true).Take(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Subscription{}, ErrCustomerNotFound
	}
	return s, err
}

// UpdateSubscription is a ...
func (s *Subscription) UpdateSubscription(uid uint32) (*Subscription, error) {
	db := config.DB.Model(&Subscription{}).Where("id = ?", uid).Take(&Subscription{}).UpdateColumns(
		map[string]interface{}{
			"customer_id": s.CustomerID,
			"stripe_id":   s.StripeID,
			"tariff_id":   s.TariffID,
			"status":      s.Status,
			"update_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Subscription{}, db.Error
	}

	err := db.Model(&Subscription{}).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	return s, nil
}

// DeleteSubscription is a ...
func (s *Subscription) DeleteSubscription(uid uint32) (int64, error) {
	db := config.DB.Model(&Subscription{}).Where("id = ?", uid).Take(&Subscription{}).Delete(&Subscription{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// CustomerSubscriptionsList is a ...
type CustomerSubscriptionsList struct {
	ID           uint32    `json:"id"`
	StripeID     string    `json:"stripe_id"`
	CustomerID   string    `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	TariffID     uint32    `json:"tariff_id"`
	TariffName   string    `json:"tariff_name"`
	Status       bool      `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// SubscriptionsByCustomerID is a ...
func SubscriptionsByCustomerID(customer_id string) *[]CustomerSubscriptionsList {
	result := []CustomerSubscriptionsList{}
	db := config.DB.Raw("SELECT subscriptions.ID,subscriptions.stripe_id,subscriptions.customer_id,customers.NAME AS customer_name,subscriptions.tariff_id,tariffs.NAME AS tariff_name,subscriptions.status,subscriptions.created_at,subscriptions.updated_at FROM subscriptions INNER JOIN customers ON subscriptions.customer_id=customers.ID INNER JOIN tariffs ON subscriptions.tariff_id=tariffs.ID WHERE subscriptions.customer_id=? ORDER BY subscriptions.created_at DESC ", customer_id).Scan(&result)
	if db.Error != nil {
		return &result
	}
	return &result
}
