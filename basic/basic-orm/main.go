package main

import (
	"projects-golang/basic/basic-orm/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router.StartServer()
}
