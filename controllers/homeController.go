package controllers

import (
	"net/http"
	"user-service/helper"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	helper.RESPONSE(w, http.StatusOK, "Welcome To This Awesome API")
}
