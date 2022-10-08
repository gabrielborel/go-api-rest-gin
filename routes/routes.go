package routes

import (
	studentsController "api-rest-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/students", studentsController.HandleGetAllStudents)
	r.POST("/students", studentsController.HandleCreateNewStudent)
	r.GET("/students/:id", studentsController.HandleGetStudentById)
	r.GET("/students/cpf/:cpf", studentsController.HandleGetStudentByCPF)
	r.DELETE("/students/:id", studentsController.HandleDeleteStudent)
	r.PATCH("/students/:id", studentsController.HandleUpdateStudent)

	r.Run()
}
