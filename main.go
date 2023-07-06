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

	r.GET("/api/questgiver/all", controllers.GetQuestgivers)
	r.GET("/api/questgiver/qualities/:npc", controllers.GetQuestgiverQualities)
	r.GET("/api/questgiver/quests/:npc", controllers.GetQuestgiverQuests)
	r.GET("/api/questgiver/qualitiesquests/:npc", controllers.GetQuestgiverQualitiesQuests)

	r.GET("/api/questgiver/generic", controllers.GetQgQlGeneric) //refactor
	r.GET("/api/questgiver/:npc", controllers.GetQuestgiver)

	r.GET("/api/quality/all", controllers.GetQualities)
	r.GET("/api/quality/:quality", controllers.GetQualityQuestgivers)

	//r.Get("/api/questgiver/quests/:npc") - instead of "/api/getall/qgname"

	r.Run()

}
