package models

import (
	"user-service/entities"
	"user-service/middlewares"

	"github.com/jinzhu/gorm"
)

type StdProfile entities.StdProfile

func (p *StdProfile) SaveProfile(db *gorm.DB) (*StdProfile, error) {
	var err error

	err = db.Debug().Create(&p).Error
	if err != nil {
		return &StdProfile{}, err
	}
	return p, nil
}

func (p *StdProfile) FindProfileByID(db *gorm.DB, ID string, middlewares *middlewares.Access) (*[]StdProfile, error) {
	var err error
	data := []StdProfile{}

	err = db.Debug().Preload("StdUniversity").Model(&StdProfile{}).Where("uid = ?", ID).Find(&data).Error
	if err != nil {
		return &[]StdProfile{}, err
	}
	return &data, nil
}

func (p *StdProfile) SaveUpdateProfile(db *gorm.DB, ID string, middlewares *middlewares.Access) (*StdProfile, error) {
	err := db.Debug().Model(&StdProfile{}).Where("uid = ?", ID).Update(&p).Error
	if err != nil {
		return &StdProfile{}, err
	}
	return p, nil
}
