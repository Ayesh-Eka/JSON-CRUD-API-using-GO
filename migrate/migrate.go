package main

import (
	"github.com/ayesh2/go-crud2/initializers"
	"github.com/ayesh2/go-crud2/models"
)

func init() {
	initializers.LoadCnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}