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
	r.POST("/api/", controllers.PostCreateFromJson)
	r.GET("/api/getall", controllers.GetAll)
	r.GET("/api/getall/qgname", controllers.GetAllQ)
	r.GET("/api/getdescriptions", controllers.GetDescriptions)
	r.GET("/api/getdescription", controllers.GetDescriptions)
	r.GET("/api/getstructures", controllers.GetStructures)
	r.GET("/api/getstructure", controllers.GetStructure)
	r.Run()

	/*http.HandleFunc("/api/", controllers.HandleJson)

	log.Fatal(http.ListenAndServe(":5000", nil))*/
}
