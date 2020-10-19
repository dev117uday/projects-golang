package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/", "./index.html")
	fmt.Println("running at : http://localhost:8080")
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
