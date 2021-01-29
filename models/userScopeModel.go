package models

import (
	"user-service/entities"
	"user-service/middlewares"

	"github.com/jinzhu/gorm"
)

type StdUserScope entities.StdUserScope

func (u *StdUserScope) SaveUserScope(db *gorm.DB, value string, tipe string) (*StdUserScope, error) {
	var scope string
	var uid string

	if tipe == "user" {
		rows, err := db.Debug().Model(&StdScope{}).Select("scope").Rows()
		for rows.Next() {
			rows.Scan(&scope)

			sQuery := `INSERT INTO std_user_scopes (scope, uid, is_active, created_at, updated_at)
			VALUES('` + scope + `', '` + value + `', false, NOW(), NOW())`
			err = db.Debug().Exec(sQuery).Error
			if err != nil {
				return &StdUserScope{}, err
			}
		}
	} else {
		rows, err := db.Debug().Model(&StdUser{}).Select("uid").Rows()
		for rows.Next() {
			rows.Scan(&uid)

			sQuery := `INSERT INTO std_user_scopes (scope, uid, is_active, created_at, updated_at)
			VALUES('` + value + `', '` + uid + `', false, NOW(), NOW())`
			err = db.Debug().Exec(sQuery).Error
			if err != nil {
				return &StdUserScope{}, err
			}
		}
	}

	return u, nil
}

func (u *StdUserScope) FindAllUserScope(db *gorm.DB, middlewares *middlewares.Access) (*[]StdUserScope, error) {
	var err error
	var data []StdUserScope

	err = db.Debug().Where("is_deleted != ?", true).Find(&data).Error
	if err != nil {
		return &[]StdUserScope{}, err
	}
	return &data, err
}

func (u *StdUserScope) FindUserScopeByID(db *gorm.DB, ID string, middlewares *middlewares.Access) (*[]StdUserScope, error) {
	var err error
	data := []StdUserScope{}

	err = db.Debug().Model(&StdUserScope{}).Where("uid = ? and is_deleted != ?", ID, true).Find(&data).Error
	if err != nil {
		return &[]StdUserScope{}, err
	}
	return &data, nil
}

func (u *StdUserScope) SaveUpdateUserScope(db *gorm.DB, ID string, middlewares *middlewares.Access) (*StdUserScope, error) {
	err := db.Debug().Model(&StdUserScope{}).Where("uid = ?", ID).Update(&u).Error
	if err != nil {
		return &StdUserScope{}, err
	}
	return u, nil
}
