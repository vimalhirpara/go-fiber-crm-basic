package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
	"github.com/vimal/go-fiber-crm-basic/database"
	"github.com/vimal/go-fiber-crm-basic/lead"
	"gorm.io/gorm"
	//_ "gorm.io/driver/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase() {
	var err error
	//database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	database.DBConn, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}
