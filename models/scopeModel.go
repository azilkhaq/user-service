package models

import (
	"user-service/entities"
	"user-service/middlewares"

	"github.com/jinzhu/gorm"
)

type StdScope entities.StdScope

func (s *StdScope) SaveScope(db *gorm.DB, middlewares *middlewares.Access) (*StdScope, error) {
	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &StdScope{}, err
	}
	return s, nil
}

func (s *StdScope) FindAllScope(db *gorm.DB, middlewares *middlewares.Access) (*[]StdScope, error) {
	var err error
	var data []StdScope

	err = db.Debug().Where("is_deleted != ?", true).Find(&data).Error
	if err != nil {
		return &[]StdScope{}, err
	}
	return &data, err
}

func (s *StdScope) FindScopeByID(db *gorm.DB, ID string, middlewares *middlewares.Access) (*[]StdScope, error) {
	var err error
	data := []StdScope{}

	err = db.Debug().Model(&StdScope{}).Where("id = ? and is_deleted != ?", ID, true).Find(&data).Error
	if err != nil {
		return &[]StdScope{}, err
	}
	return &data, nil
}

func (s *StdScope) SaveUpdateScope(db *gorm.DB, ID string, middlewares *middlewares.Access) (*StdScope, error) {
	err := db.Debug().Model(&StdScope{}).Where("id = ?", ID).Update(&s).Error
	if err != nil {
		return &StdScope{}, err
	}
	return s, nil
}

func (s *StdScope) SaveDeleteScope(db *gorm.DB, scope string, middlewares *middlewares.Access) (*StdScope, error) {
	var err error
	sQuery := `DELETE FROM std_user_scopes WHERE scope = ?`
	err = db.Debug().Exec(sQuery, scope).Error
	if err != nil {
		return &StdScope{}, err
	}

	err = db.Debug().Model(&StdScope{}).Where("scope = ?", scope).Delete(&s).Error
	if err != nil {
		return &StdScope{}, err
	}

	return s, nil
}
