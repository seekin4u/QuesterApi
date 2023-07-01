package main

import (
	"QuesterApi/controllers"
	"QuesterApi/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("App started")
	r := gin.Default()
	r.GET("/api/createdef", controllers.PostCreate)
	r.POST("/api/create", controllers.PostCreateFromJson)
	r.GET("/api/getall", controllers.GetAll)
	//r.GET("/api/getall/qgname", controllers.GetAllQ)
	r.GET("/api/getdescriptions", controllers.GetDescriptions)
	r.GET("/api/getdescription", controllers.GetDescription)
	r.GET("/api/getstructures", controllers.GetStructures)
	r.GET("/api/getstructure", controllers.GetStructure)

	r.GET("/api/questgiver/questgivers", controllers.GetQgs)
	r.GET("/api/questgiver/questgiverqualities/:npc", controllers.GetQgQs)
	r.GET("/api/questgiver/generic", controllers.GetQgQlGeneric) //refactor
	r.GET("/api/questgiver/:npc", controllers.GetQgSpecial)

	//r.Get("/api/questgiver/quests/:npc") - instead of "/api/getall/qgname"

	r.Run()

}
