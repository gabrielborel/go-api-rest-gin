package routes

import (
	"api-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

    r.LoadHTMLGlob("templates/*")
    r.Static("/assets", "./assets")

	r.GET("/students", controllers.HandleGetAllStudents)
	r.POST("/students", controllers.HandleCreateNewStudent)
	r.GET("/students/:id", controllers.HandleGetStudentById)
	r.GET("/students/cpf/:cpf", controllers.HandleGetStudentByCPF)
	r.DELETE("/students/:id", controllers.HandleDeleteStudent)
	r.PATCH("/students/:id", controllers.HandleUpdateStudent)
    r.GET("/index", controllers.HandleShowIndexPage)

    r.NoRoute(controllers.HandleShowRouteNotFoundPage)

	r.Run()
}
