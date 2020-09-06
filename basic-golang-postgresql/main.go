package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "uday"
	password = "test123"
	dbname   = "test"
)


type sandbox struct {
	id        int
	firstName string
	lastName  string
	email     string
	gender    string
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	app := fiber.New()

	app.Get("/data", func(context *fiber.Ctx) {

		start := time.Now()

		sqlStatement := "SELECT * FROM test LIMIT 4"
		data, err := db.Query(sqlStatement)

		if err != nil {
			panic(err)
		}

		snbs := make([]sandbox, 0)

		for data.Next() {
			snb := sandbox{}
			err := data.Scan(&snb.id, &snb.firstName, &snb.lastName, &snb.email, &snb.gender)
			if err != nil {
				log.Println(err)
			}
			snbs = append(snbs, snb)
		}

		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed)
		context.Send(snbs)
	})

	_ = app.Listen(3000)
}