package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data 	interface{} `json:"data"`
}

type Book struct {
	Id int 	`json:"id"`
	Title string 	`json:"title"`
}

func main(){
	e := echo.New()
	e.GET("/books", GetBooksController)
	e.GET("/books/:id", GetDetailBookController)
	e.POST("/books", AddBookController)
	e.Start(":8080")
}

func AddBookController(c echo.Context) error {
	var requestBook Book
	c.Bind(&requestBook)

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: requestBook,
	})
}

func GetDetailBookController(c echo.Context) error {
	
	id, _ := strconv.Atoi(c.Param("id"))
	
	var book Book = Book{id, "sasa"}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: book,
	})
}

func GetBooksController(c echo.Context) error {

	negara := c.QueryParam("negara")

	var books []Book
	books = append(books, Book{1, negara})

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: books,
	})
}