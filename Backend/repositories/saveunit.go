package repositories

import (
	config "backend/config"
	"backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Getsaveunitreq struct {
	Datestart *string `json:"datestart"`
	Dateend   *string  `json:"dateend"`
}

func GetallSaveunit(r Getsaveunitreq) ([]models.Saveunit, error) {
	var Saveunit []models.Saveunit
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Raw(`SELECT
	DATE_TRUNC( 'month', created_at ) AS DATE,
	SUM ( unit ) AS unit 
FROM
	saveunits 
WHERE
	created_at BETWEEN ?
	AND ?
GROUP BY
	DATE_TRUNC( 'month', created_at )`, r.Datestart, r.Dateend).Find(&Saveunit).Error; err != nil {
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
