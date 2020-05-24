package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/config"
)

type Subscription struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	StripeID   string    `gorm:"size:18;not null;unique" json:"stripe_id"`
	CustomerID uint32    `sql:"type:int REFERENCES customers(id)" json:"customer_id"`
	TarifID    uint32    `sql:"type:int REFERENCES tarifs(id)" json:"tarif_id"`
	Status     bool      `gorm:"false" json:"status"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Subscription) SaveSubscription() (*Subscription, error) {
	err := config.DB.Debug().Create(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	return s, nil
}

func (s *Subscription) FindSubscriptionByID(uid uint32) (*Subscription, error) {
	err := config.DB.Debug().Model(Subscription{}).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Subscription{}, ErrTarifNotFound
	}
	return s, err
}

func (s *Subscription) FindSubscriptionByStripeID(stripe_id string) (*Subscription, error) {
	err := config.DB.Debug().Model(Subscription{}).Where("stripe_id = ? AND status = ?", stripe_id, true).Take(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Subscription{}, ErrCustomerNotFound
	}
	return s, err
}

func (s *Subscription) UpdateSubscription(uid uint32) (*Subscription, error) {
	db := config.DB.Debug().Model(&Subscription{}).Where("id = ?", uid).Take(&Subscription{}).UpdateColumns(
		map[string]interface{}{
			"customer_id": s.CustomerID,
			"stripe_id":   s.StripeID,
			"tarif_id":    s.TarifID,
			"status":      s.Status,
			"update_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Subscription{}, db.Error
	}

	err := db.Debug().Model(&Subscription{}).Where("id = ?", uid).Take(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	return s, nil
}

func (s *Subscription) DeleteSubscription(uid uint32) (int64, error) {
	db := config.DB.Debug().Model(&Subscription{}).Where("id = ?", uid).Take(&Subscription{}).Delete(&Subscription{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
