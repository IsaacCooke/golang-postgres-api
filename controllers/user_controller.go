package controllers

import (
	"api/data"
	"api/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func getUsers(c *gin.Context) {
	db := data.SetupDB()

	rows, err := db.Query("SELECT * FROM users")

	checkError(err)

	var users []models.User

	for rows.Next() {
		var id int
		var age int
		var first_name string
		var last_name string
		var email string

		err = rows.Scan(&id, &age, &first_name, &last_name, &email)

		checkError(err)

		users = append(users, models.User{
			Id:        id,
			Age:       age,
			FirstName: first_name,
			LastName:  last_name,
			Email:     email,
		})
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	paramId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := fmt.Sprintf("recived invalid path param which is not string: %v", c.Param("id"))
		log.Println(e)
		c.IndentedJSON(http.StatusNotFound, e)
		return
	}
	db := data.SetupDB()

	var user models.User

	var id, age int
	var first_name, last_name, email string

	sqlStatement :=
		`SELECT * FROM users
		WHERE id = $1;`

	row := db.QueryRow(sqlStatement, paramId)

	err = row.Scan(&id, &age, &first_name, &last_name, &email)

	switch err {
	case sql.ErrNoRows:
		log.Printf("No rows returned with id %i", id)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	case nil:
		user = models.User{
			Id:        id,
			Age:       age,
			FirstName: first_name,
			LastName:  last_name,
			Email:     email,
		}
		c.IndentedJSON(http.StatusOK, user)
	default:
		e := fmt.Sprintf("error: %v occured while reading the database for the user record", err)
		log.Println(e)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
}

func createUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	sqlStatement :=
		`INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	err := data.Connect().QueryRow(sqlStatement, newUser.Age, newUser.Email, newUser.FirstName, newUser.LastName).Scan(&id)
	checkError(err)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	var updatedUser models.User

	if err := c.BindJSON(&updatedUser); err != nil {
		return
	}

	sqlStatement :=
		`UPDATE users
		SET first_name $2, last_name $3, email $4
		WHERE id = $1;`

	res, err := data.Connect().Exec(sqlStatement, updatedUser.Id, updatedUser.FirstName, updatedUser.LastName, updatedUser.Email)
	checkError(err)
	count, err := res.RowsAffected()
	checkError(err)

	fmt.Println(count)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	sqlStatement :=
		`DELETE FROM users
		WHERE id = $1`

	res, err := data.Connect().Exec(sqlStatement, id)
	checkError(err)
	count, err := res.RowsAffected()

	fmt.Println(count)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
