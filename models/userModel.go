package models

import (
	"errors"
	"strings"
	"time"
	"user-service/entities"
	"user-service/helper"
	"user-service/middlewares"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type StdUser entities.StdUser

const status string = "deleted"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *StdUser) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *StdUser) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Role == "" {
			return errors.New("Required Role")
		}
		if u.EmailAddress == "" {
			return errors.New("Required Email")
		}
		if u.PhoneNumber == "" {
			return errors.New("Required Phone Number")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.EmailAddress == "" {
			return errors.New("Required Email")
		}
		return nil

	default:
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.EmailAddress == "" {
			return errors.New("Required Email")
		}
		return nil
	}
}

func (u *StdUser) SaveUsers(db *gorm.DB) (*StdUser, error) {
	var err error
	u.Uid = helper.GENERATEUUID()
	u.Status = "enabled"

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &StdUser{}, err
	}
	return u, nil
}

func (u *StdUser) FindAllUsers(db *gorm.DB, middlewares *middlewares.Access) (*[]StdUser, error) {
	var err error
	var users []StdUser

	err = db.Debug().Where("status != ?", status).Find(&users).Error
	if err != nil {
		return &[]StdUser{}, err
	}
	return &users, err
}

func (u *StdUser) FindUsersByID(db *gorm.DB, uid string, middlewares *middlewares.Access) (*[]StdUser, error) {
	var err error
	users := []StdUser{}

	err = db.Debug().Model(&StdUser{}).Where("uid = ? and status != ?", uid, status).Find(&users).Error
	if err != nil {
		return &[]StdUser{}, err
	}
	return &users, nil
}

func (u *StdUser) SaveUpdateUsers(db *gorm.DB, uid string, middlewares *middlewares.Access) (*StdUser, error) {
	u.UpdatedAt = time.Now()
	err := db.Debug().Model(&StdUser{}).Where("uid = ?", uid).Update(&u).Error
	if err != nil {
		return &StdUser{}, err
	}
	return u, nil
}

func (u *StdUser) SaveDeleteUsers(db *gorm.DB, uid string, middlewares *middlewares.Access) (*StdUser, error) {
	u.Status = "deleted"
	err := db.Debug().Model(&StdUser{}).Where("uid = ?", uid).Update(&u).Error
	if err != nil {
		return &StdUser{}, err
	}
	return u, nil
}
