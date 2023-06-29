package main

import (
	"QuesterApi/initializers"
	"QuesterApi/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.QuestTime{}, &models.QuestStructure{}, &models.QuestDescription{})
}
