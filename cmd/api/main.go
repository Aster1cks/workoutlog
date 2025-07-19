package main

//Now in branch GIN

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Aster1cks/workoutlog/internal/database"
	"github.com/Aster1cks/workoutlog/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load env %v", err)
	}
	url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Couldn't open db %v", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Couldn't ping db %v", err)
	}

	fmt.Println("Db ping succesfull")

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infolog.Printf("DB Ping succesfull")
	var app = &server.Application{
		InfoLogger:  infolog,
		ErrorLogger: errorlog,
		Workouts: &database.WorkoutModel{
			DB: db,
		},
	}

	// var srv = &http.Server{
	// 	Addr:     os.Getenv("PORT"),
	// 	ErrorLog: errorlog,
	// 	Handler:  app.Routes(),
	// }
	addr := os.Getenv("PORT")

	ginsrv := app.Routes()
	infolog.Printf("Starting server on %s", addr)
	// err = http.ListenAndServe(srv.Addr, srv.Handler)
	err = ginsrv.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
