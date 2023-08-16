package main

import (
	"os"
	"prakerja8/configs"
	"prakerja8/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main(){
	// loadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	
	e.Start(getPort())
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
	  return ":" + envPort
	}
	return ":8080"
  }
  

func loadEnv(){
	err := godotenv.Load()
  	if err != nil {
   	 	panic("Error loading .env file")
  	}
}