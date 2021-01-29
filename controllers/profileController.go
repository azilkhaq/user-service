package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"user-service/helper"
	"user-service/middlewares"
	"user-service/models"

	"github.com/gorilla/mux"
)

func (server *Server) CreateProfile(w http.ResponseWriter, r *http.Request) {
	data := &models.StdProfile{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := data.SaveProfile(server.DB)
	if err != nil {
		format := helper.FormatError(err.Error())
		helper.ERROR(w, http.StatusUnprocessableEntity, format)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, result.ID))
	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusCreated,
		M{
			"data":    result,
			"message": "Successfully",
			"status":  http.StatusOK,
		},
	)
}

func (server *Server) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := models.StdProfile{}
	result, err := data.FindProfileByID(server.DB, ID, token)
	if err != nil {
		helper.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK,
		M{
			"data":    result,
			"status":  http.StatusOK,
			"message": "Successfully",
		},
	)
}

func (server *Server) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token, err := middlewares.ExtractTokenMetadata(r)
	if err != nil {
		helper.ERROR(w, http.StatusUnauthorized, errors.New("Expired token"))
		return
	}

	vars := mux.Vars(r)
	ID := vars["id"]

	data := &models.StdProfile{}
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = data.SaveUpdateProfile(server.DB, ID, token)
	if err != nil {
		helper.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	type M map[string]interface{}
	helper.RESPONSE(w, http.StatusOK, M{"status": http.StatusOK, "message": "Successfully"})
}
