package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT         string = ":80"
	MYSQL_DB     string = "/claire"
	MYSQL_DBUSER string = "root"
	MYSQL_DBPASS string = "" // eg ":mypasswd"
	MYSQL_DBHOST string = "mysql"
)

func main() {
	// Instantiate our JsonRenderer for our api
	renderer := JsonRenderer{}

	// Instantiate BaseController with our JsonRenderer
	baseController := BaseController{
		IRenderer: &renderer,
	}

	// Instantiate repositories
	mysqlDb := MysqlDb{}
	mysqlDb.Init()

	makeRepo := MakeMysqlRepo{
		MysqlDb: &mysqlDb,
	}
	modelRepo := ModelMysqlRepo{
		MysqlDb: &mysqlDb,
	}
	tagRepo := TagMysqlRepo{
		MysqlDb: &mysqlDb,
	}

	defer func() {
		mysqlDb.Deinit()
	}()

	// Instantiate application controllers
	makeController := MakeController{
		BaseController: baseController,
		makeRepo:       &makeRepo,
	}
	modelController := ModelController{
		BaseController: baseController,
		modelRepo:      &modelRepo,
	}
	tagController := TagController{
		BaseController: baseController,
		tagRepo:        &tagRepo,
	}

	// Setup routes
	router := mux.NewRouter().StrictSlash(false)

	// GET /api/make
	router.HandleFunc("/api/make", makeController.Index).Methods("GET")
	// GET /api/make/{id}
	router.HandleFunc("/api/make/{id}", makeController.View).Methods("GET")

	// GET /api/model
	router.HandleFunc("/api/model", modelController.Index).Methods("GET")
	// GET /api/model/{id}
	router.HandleFunc("/api/model/{id}", modelController.View).Methods("GET")

	// GET /api/tag
	router.HandleFunc("/api/tag", tagController.Index).Methods("GET")
	// GET /api/tag/{id}
	router.HandleFunc("/api/tag/{id}", tagController.View).Methods("GET")

	log.Println("Listening on port 80...")
	http.ListenAndServe(PORT, router)
}
