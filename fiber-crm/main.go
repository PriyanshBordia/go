package main

import (
	"fmt"

	"github.com/PriyanshBordia/go/fiber-crm/lead"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

var print = fmt.Println

func routers(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.AddLead)
	// app.Put("/api/v1/lead/", lead.UpdateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)

}

func init() {
	var err error 
	database.db, err = gorm.Open("sqlite3", "leads.db")
	if err {
		print(err)
	}
	database.db.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	init()
	routers(app)
	app.Listen(3000)
	defer database.db.Close()
}