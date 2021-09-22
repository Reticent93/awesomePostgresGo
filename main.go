package main

//var DB *gorm.DB

func main() {
	//host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASS")
	//dbname := os.Getenv("DB_NAME")

	//<--------------------------CREATE-------------------------->
	//	sqlStatement := `INSERT INTO cars(make, model, mileage, clean)
	//VALUES($1, $2, $3, $4)
	//RETURNING id`
	//	id := 0
	//	err = db.QueryRow(sqlStatement, "Nissan", "Altima", 87946, false).Scan(&id)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("New record ID is: ", id)
	//

	//<-------------------------UPDATE---------------------------->
	//	sqlStatement := `
	//UPDATE cars
	//SET make = $2, model = $3
	//WHERE id = $1
	//RETURNING id, make, model`
	//
	//	res, err := db.Exec(sqlStatement, 4, "Audi", "A6")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	count, err := res.RowsAffected()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(count)
	//

	//<-------------------------DELETE------------------------------>
	//	sqlDelete := `
	//DELETE FROM cars
	//WHERE id = $1`
	//	_, err =db.Exec(sqlStatement, 7)
	//	if err != nil {
	//		panic(err)
	//	}
	//

	//<-----------------------BETTER UPDATE----------------------->
	//	sqlState :=`
	//UPDATE cars
	//SET make = $2, model = $3
	//WHERE id = $1
	//RETURNING id, make, model`
	//
	//var make string
	//	var model string
	//var id int
	//
	//err = db.QueryRow(sqlState, 4, "Acura", "NSX").Scan(&id, &make, &model)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(id, make, model)

	//<-----------------------------QUERY SINGLE ROW--------------------------->
	//sqlQuery := `SELECT id, make, model FROM cars WHERE id=$1`
	//
	//var make string
	//var model string
	//var id int
	//row := db.QueryRow(sqlQuery, 6)
	//	switch err := row.Scan(&id, &make, &model); err {
	//	case sql.ErrNoRows:
	//		fmt.Println("No Rows Returned")
	//	case nil:
	//		fmt.Println(id, make, model)
	//	default:
	//		panic(err)
	//
	//	}

	//<-----------------QUERY MANY RECORDS--------------------------------------->
	//	rows, err := db.Query("SELECT id, make, model FROM cars LIMIT $1", 7)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer rows.Close()
	//	for rows.Next() {
	//		var id int
	//		var make string
	//		var model string
	//		err = rows.Scan(&id, &make, &model)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		fmt.Println(id, make, model)
	//	}
	//
	//	err = rows.Err()
	//	if err != nil {
	//		panic(err)
	//	}
	//
}
