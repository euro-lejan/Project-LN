package repositories

import (
	config "backend/config"
	"backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetallSaveunit(r string) ([]models.Saveunit, error) {
	var Saveunit []models.Saveunit
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Where("date LIKE ?", r+"%").Order("id").Find(&Saveunit).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Saveunit, err
	}

	tx.Commit()

	return Saveunit, nil
}

func CreateSaveunit(d models.Saveunit) (models.Saveunit, error) {
	tx := config.DB.Begin()
	if err := tx.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}

func UpdateSaveunit(d models.Saveunit) (models.Saveunit, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}
