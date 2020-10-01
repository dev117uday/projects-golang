package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// "gorm.io/driver/sqlite"
// _ "gorm.io/gorm"
// "github.com/jinzhu/gorm/dialects/sqlite"
// "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/sqlite"

var db *gorm.DB
var err error

// User struct : to interact with the database
type User struct {
	gorm.Model
	Name  string
	Email string
}

// InitialMigration  :  to create a connection
func InitialMigration() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Panic")
	}
	db.Close()
}

// GetAllUsers : to get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// CreateNewUser : to create new user
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "createNewUser endpoint hit")
}

// DeleteUser : to delete user from db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deteleUser endpoint hit")
}

// UpdateUser :  to update user info
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUser endpoint hit")
}
