package main

import (
	"github.com/shamsher4499/go-crud/initializers"
	"github.com/shamsher4499/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

