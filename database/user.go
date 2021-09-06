package database

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "me:password@tcp(127.0.0.1:3306)/tester"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		print("Cannot connect to server")
	}

	DB.AutoMigrate()
}

func setJsonHeader(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	setJsonHeader(res)
	var users []User
	DB.Find(&users)
	json.NewEncoder(res).Encode(users)
}
func GetUser(res http.ResponseWriter, req *http.Request) {
	setJsonHeader(res)
	var user User
	params := mux.Vars(req)
	DB.First(&user, params["id"])
	json.NewEncoder(res).Encode(user)
}
func CreateUser(res http.ResponseWriter, req *http.Request) {
	setJsonHeader(res)
	var user User
	json.NewDecoder(req.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(res).Encode(user)
}
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	setJsonHeader(res)
	params := mux.Vars(req)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(req.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(res).Encode(user)
}
func DeleteUsers(res http.ResponseWriter, req *http.Request) {
	setJsonHeader(res)
	params := mux.Vars(req)
	var user User
	DB.Delete(&user, params["id"])
	text := fmt.Sprintf("The user with ID %s was deleted", params["id"])
	json.NewEncoder(res).Encode(text)
}
