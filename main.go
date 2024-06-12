package main

import (
	"database/sql"
	"encoding/json"
	"golang-sqlite-crud/models"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	createTable := `
		CREATE TABLE IF NOT EXISTS employee (id INTEGER NOT NULL PRIMARY KEY, name TEXT, jobName TEXT);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName: "Golang SQLite CRUD",
	})

	app.Post("/employee", func(c *fiber.Ctx) error {

		payload := new(models.Employee)

		if err := c.BodyParser(payload); err != nil {
			return err
		}

		db.Exec("INSERT INTO employee ( name, jobName ) VALUES (?,?)", payload.Name, payload.JobName)

		return c.SendStatus(fiber.StatusAccepted)
	})

	app.Get("/employees", func(c *fiber.Ctx) error {
		var result []models.Employee = []models.Employee{}

		rows, err := db.Query("SELECT * FROM employee")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name string
			var jobName string
			err = rows.Scan(&id, &name, &jobName)

			if err != nil {
				log.Fatal(err)
			}
			result = append(result, models.Employee{
				Id:      id,
				Name:    name,
				JobName: jobName,
			})
		}

		bytes, _ := json.Marshal(result)

		return c.Send(bytes)

	})

	app.Get("/employee/:id", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")

		var (
			result   models.Response
			employee models.Employee
			id       string = c.Params("id")
			message  string = "NOT FOUND"
		)

		row, err := db.Query("SELECT * FROM employee WHERE id = ?", id)
		if err != nil {
			log.Fatal(err.Error())
		}

		for row.Next() {
			message = "FOUND"
			var id int
			var name string
			var jobName string
			err = row.Scan(&id, &name, &jobName)

			if err != nil {
				log.Fatal(err)
			}
			employee = models.Employee{
				Id:      id,
				Name:    name,
				JobName: jobName,
			}
		}

		result = models.Response{
			Message:  message,
			Employee: employee,
		}

		bytes, _ := json.Marshal(result)

		return c.Send(bytes)
	})

	log.Fatal(app.Listen(":3000"))
}
