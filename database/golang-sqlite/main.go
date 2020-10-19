package main

import (
	"database/sql"
	"fmt"
	"log"
	"projects-golang/golang-sqlite/newsfeed"

	"github.com/gofiber/fiber"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./newsfeed.db")
	if err != nil {
		fmt.Println("PROBLEM")
		panic(err.Error())
	}
	feed := newsfeed.NewFeed(db)

	app := fiber.New()

	app.Get("/post", func(con *fiber.Ctx) {
		items := feed.GetItem()
		con.JSON(items)
	})

	app.Post("/post", func(con *fiber.Ctx) {
		p := new(newsfeed.Item)

		if err := con.QueryParser(p); err != nil {
			panic(err.Error())
		}
		feed.AddItem(newsfeed.Item{
			Content: p.Content,
		})
		con.Send("success")
	})

	fmt.Println("Running at : http://localhost:8000")
	log.Fatal(app.Listen(":8000"))

}
