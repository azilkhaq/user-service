package controllers

import (
	"fmt"
	"log"
	"net/http"
	"user-service/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	var DBURL string

	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
		log.Fatal("This is the error:", err)
	}

	server.DB.Debug().AutoMigrate(&models.StdUser{}, &models.StdProfile{},
		&models.StdScope{}, &models.StdUserScope{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to " + addr)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(server.Router)))
}
