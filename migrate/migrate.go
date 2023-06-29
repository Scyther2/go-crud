package main

import (
	"github.com/Scyther2/go-crud/initializers"
	"github.com/Scyther2/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
