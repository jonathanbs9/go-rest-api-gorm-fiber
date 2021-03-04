package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jonathanbs9/go-rest-api-gorm-fiber/book"
	"github.com/jonathanbs9/go-rest-api-gorm-fiber/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloFunc(c *fiber.Ctx) {
	c.Send("Hello Golang! ")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBookById)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Fallo la conexion a la base de datos")
	}
	fmt.Println("Conexión exitosa a la base de datos! ")

	// Automigration => take the struct defined and create a table
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated!")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)
	app.Listen(3000)
	fmt.Println("Serving on port: 3000")
}
