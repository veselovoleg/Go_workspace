package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/go-martini/martini"
)

var db *sql.DB

func main() {

	m := martini.Classic()
	m.Get("/", messageFromDB)
	m.Get("/answer", messageToDB)
	m.Run()

}

func messageFromDB() {
	db = dbConnect()
	//Get data from DB
	rowsContent, err := db.Query("SELECT * FROM requests") //Receive data from content table
	if err != nil {
		panic(err)
	}

	contFromDB := []database{}
	contFromDBtemp := database{}

	for count := 0; rowsContent.Next(); count++ {
		rowsContent.Scan(&contFromDBtemp.Reqnum, &contFromDBtemp.Reqtext)
		contFromDB = append(contFromDB, contFromDBtemp)
	}

	fmt.Println(contFromDB)

}

func messageToDB() {
	db = dbConnect()
	//Get data from DB

	SendDataStruct := database{
		Reqtext: "Third",
	}

	_, err := db.Exec("INSERT INTO requests(reqtext) VALUES ($1) ON CONFLICT DO NOTHING;",
		SendDataStruct.Reqtext)
	if err != nil {
		panic(err)
	}

}

func dbConnect() *sql.DB {
	//Connect to DB
	connStr := "user=oleg password=p dbname=siteexample sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")
	return db
}

type database struct {
	Reqnum  int
	Reqtext string
}
