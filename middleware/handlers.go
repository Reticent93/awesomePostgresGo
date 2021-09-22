package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/Reticent93/awesomePostgresGo/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

//response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//create function with postgres
func create() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	//Opening connection to database
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		if err != nil {
			panic(err)
		}
	}(db)

	fmt.Println("Successfully connected")
	return db
}

//CreateUser create a car in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//setting the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	//create an empty car of type models.Cars
	var cars models.Cars

	//decode the json request to car
	err := json.NewDecoder(r.Body).Decode(cars)
	if err != nil {
		log.Fatal("Unable to decode the request body", err)
	}

	//call insert car function and pass in a car
	insertID := insertCars(cars)

}
