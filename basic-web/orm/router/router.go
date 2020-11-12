package router

import (
	"fmt"
	"log"
	"net/http"
	"projects-golang/basic/basic-orm/database"

	"github.com/gofiber/fiber/v2"
)

//StartServer : contains all routing information
func StartServer() {

	database.CreateConnection()

	app := fiber.New()
	fmt.Println("Running at : http://localhost:8000")

	app.Get("/", func(con *fiber.Ctx) error {
		return con.SendString("Hello World")
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
		} else if result == -1 {
			return con.SendString("user doesnt exist in DB")
		}
		return con.SendString("user delete")

	})
	// UpdateUser :  to update user info
	app.Put("/update/:name/:email", func(con *fiber.Ctx) error {
		_, error := database.UpdateUser(con.Params("name"), con.Params("email"))
		if error != nil {
			con.SendString("error in updating the user")
		}
		return con.SendString("user updated")
	})

	log.Fatal(app.Listen(":8000"))
}

func helloWorld(writer http.ResponseWriter, reader *http.Request) {
	data := []byte("Hello World")
	writer.Write(data)
}
