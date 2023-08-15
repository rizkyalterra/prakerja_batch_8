package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type User struct {
	Id int 					`json:"id"` 
	Name string 			`json:"name"`
	CreditCard CreditCard
}
  
type CreditCard struct {
	Id int 				`json:"id"`
	Number string       `json:"number"`
	UserId int
}

func main(){
	initDatabase()
	e := echo.New()
	e.GET("/books", GetBooksController)
	e.GET("/users", GetUsersController)
	e.GET("/books/:id", GetDetailBookController)
	e.POST("/books", AddBookController)
	e.Start(":8080")
}

var DB *gorm.DB

func initDatabase(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja8?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

func initMigration(){
	DB.AutoMigrate(&Book{}, &User{}, &CreditCard{})
}

func AddBookController(c echo.Context) error {
	var requestBook Book
	c.Bind(&requestBook)

	// masukkan ke database
	result := DB.Create(&requestBook)
	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status: false,
			Message: "Failed insert data books",
			Data: nil,
		})
	}

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

	var books []Book
	
	result := DB.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status: false,
			Message: "Failed insert data books",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: books,
	})
}

func GetUsersController(c echo.Context) error {

	var users []User
	
	result := DB.Preload("CreditCards").Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status: false,
			Message: "Failed insert data books",
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Berhasil",
		Data: users,
	})
}