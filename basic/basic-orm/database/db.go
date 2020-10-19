package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

//UserModel : model of user table in db
type UserModel struct {
	ID    int
	Name  string
	Email string
}

// CreateConnection : to connnection to db
func CreateConnection() {
	dbx, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		fmt.Println("PROBLEM in createConnection")
		panic(err.Error())
	}
	db = dbx
	stmnt, err := db.Prepare(`
		create table if not exists users
			(
				ID integer not null
					constraint newsfeed_pk
						primary key autoincrement,
				Name text,
				Email text unique
			);
	`)
	if err != nil {
		fmt.Println("error in createconnection")
		panic(err.Error())
	}
	stmnt.Exec()
}

// GetAllUsers : to get all users from the database
func GetAllUsers() []UserModel {
	var queryString = "select * from users;"
	rows, err := db.Query(queryString)
	if err != nil {
		fmt.Println("Error executing getAll statement")
		panic(err.Error())
	}
	users := []UserModel{}
	var id int
	var name, email string
	for rows.Next() {
		rows.Scan(&id, &name, &email)
		payload := UserModel{
			ID:    id,
			Name:  name,
			Email: email,
		}
		users = append(users, payload)
	}
	return users
}

//AddUser : to add user to db
func (user *UserModel) AddUser() (int, error) {
	queryToRun := "INSERT INTO users (Name, Email) VALUES ('" + string(user.Name) + "', ' " + user.Email + "');"
	stmnt, err := db.Prepare(queryToRun)
	answer, errx := stmnt.Exec()
	if err != nil {
		fmt.Println("problem is adding user to db")
		fmt.Println(answer, errx.Error())
		fmt.Println(err.Error())
		return 0, err
	}
	return 1, nil
}

// DeleteUser : to delete user from db
func DeleteUser(email string) int {

	queryToRun := "delete from users where Email like '%" + string(email) + "%'"
	stmnt, err := db.Prepare(queryToRun)
	if err != nil {
		fmt.Println("Error in delete user")
		return 0
	}
	answer, errx := stmnt.Exec()
	noOfRowsDeleted, _ := answer.RowsAffected()
	if noOfRowsDeleted == 0 {
		return -1
	}
	if errx != nil {
		fmt.Println("Error in delete user")
		return 0
	}
	return 1
}

// UpdateUser :  to update the user information
func UpdateUser(name, email string) (int, error) {
	queryToRun := "update users set Email = '" + email + "' where Name = '" + name + "';"
	stmnt, err := db.Prepare(queryToRun)
	if err != nil {
		fmt.Println("error in updateUser")
		return 0, err
	}
	_, errx := stmnt.Exec()
	if errx != nil {
		fmt.Println("error in updateUser")
		return 0, err
	}
	return 1, nil
}
