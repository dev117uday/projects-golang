package router

import (
	"fmt"
	"log"
	"projects-golang/basic-orm/database"

	"github.com/gofiber/fiber/v2"
)

//StartServer : contains all routing information
func StartServer() {

	database.CreateConnection()

	app := fiber.New()
	fmt.Println("Running at : http://localhost:8000")

	// HelloWorld : route to "/"
	app.Get("/", func(con *fiber.Ctx) error {
		return con.SendString("HelloWorld endpoint hit")
	})
	// GetAllUsers : to get all users
	app.Get("/users", func(con *fiber.Ctx) error {
		result := database.GetAllUsers()
		return con.JSON(result)
	})
	// CreateNewUser : to create new user
	app.Post("/user", func(con *fiber.Ctx) error {
		user := new(database.UserModel)

		if err := con.QueryParser(user); err != nil {
			fmt.Println("Problem in post request")
			panic(err.Error())
		}
		_, err := user.AddUser()
		if err != nil {
			return con.SendString("error recording the user")
		}
		return con.SendString("success")
	})
	// DeleteUser : to delete user from db
	app.Delete("/delete/:email", func(con *fiber.Ctx) error {
		result := database.DeleteUser(con.Params("email"))
		if result == 0 {
			return con.SendString("unable to delete")
		}
		return con.SendString("user delete")

	})
	// UpdateUser :  to update user info
	app.Put("/update/:name/:email", func(con *fiber.Ctx) error {
		return con.SendString("UpdateUser endpoint hit")
	})

	log.Fatal(app.Listen(":8000"))
}
