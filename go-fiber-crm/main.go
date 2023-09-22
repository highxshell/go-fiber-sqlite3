package main

import (
	"fmt"
	"go-fiber-crm/database"
	"go-fiber-crm/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App){
	app.Get("/apu/v1/lead", lead.GetLeads)
	app.Get("/apu/v1/lead/:id", lead.GetLead)
	app.Post("/apu/v1/lead", lead.NewLead)
	app.Delete("/apu/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to DB")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}