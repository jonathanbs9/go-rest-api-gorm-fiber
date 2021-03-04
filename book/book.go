package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/jonathanbs9/go-rest-api-gorm-fiber/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetBooks func => Get All Books
func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
	//c.Send("<<< Todos los libros >>> ")
}

// GetBookById func => Get Book by Id
func GetBookById(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(id)
	c.JSON(book)
}

// NewBook func => Create Book
func NewBook(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	// Parseo los datos del body a book
	if err := c.BodyParser(book); err != nil {
		c.Status(400).Send(err.Error())
		return
	}
	// Caso exitoso, pego en base y muestro el book
	db.Create(&book)
	c.JSON(book)
}

// DeleteBook func => Delete Book by Id
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(400).Send("Book not found on DB with the given ID : " + id)
		return
	}
	db.Delete(book)
	c.Send("<<< Book deleted !!  >>> ")
}
