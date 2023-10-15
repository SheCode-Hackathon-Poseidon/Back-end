package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sample/api/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
)

// Server is...
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initalize is...
func (server *Server) Initalize(DbDriver, DbUser, DbPassword, DbPort, DbName, DbHost string) {
	var err error

	if DbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

		server.DB, err = gorm.Open(DbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)

			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DbDriver)
		}
	}

	if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)

			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the %s database\n", DbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Project{}, &models.Progress{}, &models.Comment{}, &models.Task{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()

	fmt.Printf("Initialized server ")
}

// Run is...
func (server *Server) Run(addr string) error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                   // Allow requests from any origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowOriginFunc: func(origin string) bool {
            return true
        },
	})

	handler := c.Handler(server.Router)
	fmt.Println("Starting to listen on " + addr)

	err := http.ListenAndServe(addr, handler)

	if err != nil {
		return err
	}

	return nil
}
