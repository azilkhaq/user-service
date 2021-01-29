package controllers

import (
	"net/http"
	"user-service/entities"
	"user-service/helper"
	"user-service/middlewares"
	"user-service/models"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	type M map[string]interface{}
	email, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, "Invalid email or password . ", http.StatusBadRequest)
		return
	}
	token, err := server.SignIn(email, password)
	if err != nil {
		helper.RESPONSE(w, http.StatusUnauthorized, M{
			"status":  http.StatusBadRequest,
			"message": "email or password incorrect",
		})
		return
	}

	helper.RESPONSE(w, http.StatusOK,
		M{
			"data":    token,
			"message": "Successfully",
			"status":  http.StatusOK,
		},
	)
}

func (server *Server) SignIn(email string, password string) (map[string]string, error) {
	var err error

	users := entities.StdUser{}
	err = server.DB.Debug().Model(entities.StdUser{}).Where("email_address = ?", email).Take(&users).Error
	if err != nil {
		return nil, err
	}

	err = models.VerifyPassword(users.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return middlewares.CreateToken(users.Uid, users.EmailAddress, users.PhoneNumber, users.Role)
}
