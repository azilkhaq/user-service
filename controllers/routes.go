package controllers

import "user-service/middlewares"

func (server *Server) initializeRoutes() {
	route := server.Router.HandleFunc
	setMiddleJSON := middlewares.SetMiddlewareJSON

	// Home
	route("/", setMiddleJSON(server.Home)).Methods("GET")

	// Auth
	route("/login", setMiddleJSON(server.Login)).Methods("POST")

	// Users
	route("/users/add", setMiddleJSON(server.CreateUsers)).Methods("POST")
	route("/users/get", setMiddleJSON(server.GetAllUsers)).Methods("GET")
	route("/users/get/{id}", setMiddleJSON(server.GetUsersByID)).Methods("GET")
	route("/users/update/{id}", setMiddleJSON(server.UpdateUsers)).Methods("PUT")
	route("/users/delete/{id}", setMiddleJSON(server.DeleteUsers)).Methods("DELETE")

	// Scope
	route("/scope/add", setMiddleJSON(server.CreateScope)).Methods("POST")
	route("/scope/get", setMiddleJSON(server.GetAllScope)).Methods("GET")
	route("/scope/delete/{id}", setMiddleJSON(server.DeleteScope)).Methods("DELETE")

	// Users Scope
	route("/user/scope/get", setMiddleJSON(server.GetAllUserScope)).Methods("GET")
	route("/user/scope/get/{id}", setMiddleJSON(server.GetUserScopeByID)).Methods("GET")
	route("/user/scope/update/{id}", setMiddleJSON(server.UpdateUserScope)).Methods("PUT")

	// Profile
	route("/profile/add", setMiddleJSON(server.CreateProfile)).Methods("POST")
	route("/profile/get/{id}", setMiddleJSON(server.GetProfileByID)).Methods("GET")
	route("/profile/update/{id}", setMiddleJSON(server.UpdateProfile)).Methods("PUT")
}
