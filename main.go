package main

import (
	"fmt"

	"github.com/apiorno/go-fiber-crm-basic/database"
	"github.com/apiorno/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(3000)

}

func initDatabases() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Couldn't connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConnection.AutoMigrate(&lead.Lead{})
}
func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.NewLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}
