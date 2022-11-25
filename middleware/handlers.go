package middleware

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/Reticent93/awesomePostgresGo/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create function with postgres
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

// CreateCar create a car in the postgres db
func CreateCar(w http.ResponseWriter, r *http.Request) {
	//setting the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	//create an empty car of type models.Cars
	var car models.Cars

	//decode the json request to car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatal("Unable to decode the request body", err)
	}

	//call insert car function and pass in a car
	insertID := insertCar(car)

	//format a response object
	res := response{
		ID:      insertID,
		Message: "Car created successfully",
	}

	//send the response
	json.NewEncoder(w).Encode(res)
}

// GetCar returns a single car
func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//get the carid from the request params, key is "id"
	params := mux.Vars(r)

	//convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	//call the getCar function with car id to retrieve a single car
	car, err := getCar(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	//send the response
	json.NewEncoder(w).Encode(car)
}

// GetAllCars will return all the cars
func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//get all the cars in the db
	cars, err := getAllCars()
	if err != nil {
		log.Fatalf("Unable to get all cars. %v", err)
	}

	//send all the cars as a response
	json.NewEncoder(w).Encode(cars)
}

// UpdateCar updates the cars' details in the db
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	//convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert the string into int", err)
	}

	//create an empty car of type models.Car
	var car models.Cars

	//decode the json request to car
	err = json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	//call update car to update the car
	//updatedRows := updateCar(int64(id), car)
	updatedRows := json.NewDecoder(r.Body).Decode(&car)

	//format the message string
	msg := fmt.Sprintf("Car updated successfully. Total rows/record affected %v", updatedRows)

	//format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	//send the response
	json.NewEncoder(w).Encode(res)

}

// DeleteCar deletes cars' detail in the db
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//get the carid from the request params, key is "id
	params := mux.Vars(r)

	//convert the id in string to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	//call the deleteCar, convert the int to int64
	//deletedRows := DeleteCar(id, r)
	deletedRows := json.NewDecoder(r.Body).Decode(&id)

	//format the message string
	msg := fmt.Sprintf("Car updated successfully. Total rows/record affected %v", deletedRows)

	//format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	//send the response
	json.NewEncoder(w).Encode(res)

}

//<-------------------------handler functions----------------------------->

// insert one car in the DB
func insertCar(car models.Cars) int64 {

	//create a postgres connection
	db := create()

	//create the insert sql query
	//returning carid will return the id of the inserted car
	sqlStatement := `INSERT INTO cars (carMake, carModel, mileage, clean) VALUES($1, $2, $3, $4) RETURNING carid`

	//the inserted id will be stored in this id
	var id int64

	//execute the sql statement
	//Scan function will save the insert id in the id
	err := db.First(sqlStatement, car.CarMake, car.CarModel, car.Mileage, car.Clean).Scan(id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)

	//return the inserted id
	return id
}

// get one car from the db by its carid
func getCar(id int64) (models.Cars, error) {
	//create a db connection
	db := create()

	//create a car of models.Cars type
	var car models.Cars

	//create the select sql query
	sqlStatement := `SELECT * FROM cars WHERE carid=$1`

	//execute the sql statement
	row := db.Find(sqlStatement, id)

	//unmarshal the row object to car.Make sure to add .Error at end
	err := row.Scan(&car).Error

	switch err {
	case gorm.ErrRecordNotFound:
		fmt.Println("No rows were returned")
		return car, nil
	case nil:
		return car, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	//return empty user on error
	return car, err
}

// get all cars from the db
func getAllCars() ([]models.Cars, error) {

	//create db connection
	db := create()

	var cars []models.Cars

	//create a select sql query
	sqlStatement := `SELECT * FROM cars`

	//execute the sql statement
	rows := db.Find(sqlStatement)

	//iterate over the rows
	for rows.Next() {
		var car models.Cars

		//unmarshal the row object to car
		err := rows.Scan(&car).Error
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		//append the car in the cars slice
		cars = append(cars, car)
	}
	//return empty car on error

	return cars, nil
}
