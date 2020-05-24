package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/werbot/lime/config"
)

type Tarif struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Price     int       `gorm:"size:6;not null;unique" json:"price"`
	Servers   int       `gorm:"size:6;not null;unique" json:"servers"`
	Companies int       `gorm:"size:6;not null;unique" json:"companies"`
	Users     int       `gorm:"size:6;not null;unique" json:"users"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Tarif) SaveTarif() (*Tarif, error) {
	err := config.DB.Debug().Create(&t).Error
	if err != nil {
		return &Tarif{}, err
	}
	return t, nil
}

func (t *Tarif) FindTarifByID(uid uint32) (*Tarif, error) {
	err := config.DB.Debug().Model(Tarif{}).Where("id = ?", uid).Take(&t).Error
	if err != nil {
		return &Tarif{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Tarif{}, ErrTarifNotFound
	}
	return t, err
}

func (t *Tarif) UpdateTarif(uid uint32) (*Tarif, error) {
	db := config.DB.Debug().Model(&Tarif{}).Where("id = ?", uid).Take(&Tarif{}).UpdateColumns(
		map[string]interface{}{
			"name":      t.Name,
			"price":     t.Price,
			"servers":   t.Servers,
			"companies": t.Companies,
			"users":     t.Users,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Tarif{}, db.Error
	}

	err := db.Debug().Model(&Tarif{}).Where("id = ?", uid).Take(&t).Error
	if err != nil {
		return &Tarif{}, err
	}
	return t, nil
}

func (t *Tarif) DeleteTarif(uid uint32) (int64, error) {
	db := config.DB.Debug().Model(&Tarif{}).Where("id = ?", uid).Take(&Tarif{}).Delete(&Tarif{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
