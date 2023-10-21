package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Id   int
	Name string
}

func getConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/todo")
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("cannot connect to database")
	}

	return db
}

// getDat is method to get data
func getData(c *fiber.Ctx) error {
	var conn = getConn()
	rows, err := conn.Query("select id, name from people")
	if err != nil {
		return c.JSON(err)
	}

	var result []Person = []Person{}

	for rows.Next() {
		var each Person = Person{}
		var err = rows.Scan(&each.Id, &each.Name)
		if err != nil {
			return c.JSON("parsing error")
		}
		result = append(result, each)
	}

	return c.JSON(result)
}

// postDat is method to post data
func postData(c *fiber.Ctx) error {

	return c.JSON("null")
}

// updateData is method to update data
func updateData(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(id)
}

// deleteData is method to delete data
func deleteData(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString(id)
}

func main() {

	getConn()
	app := fiber.New()
	app.Get("/", getData)
	app.Post("/", postData)
	app.Put("/:id", updateData)
	app.Delete("/:id", deleteData)

	app.Listen(":8000")
}
