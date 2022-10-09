package controllers

import (
	"api-rest-gin/database"
	"api-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetAllStudents(ctx *gin.Context) {
	var students []models.Student	

	database.DB.Find(&students)

	ctx.JSON(http.StatusOK, students) 
}

func HandleGetStudentById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, student) 
}

func HandleCreateNewStudent(ctx *gin.Context) {
	var student models.Student
	
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.ValidateStudentData(&student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	ctx.JSON(http.StatusCreated, student)
}

func HandleDeleteStudent(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var student models.Student

	database.DB.Delete(&student, id)

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Aluno deletado com sucesso!",
	})
}

func HandleUpdateStudent(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var student models.Student
	
	database.DB.First(&student, id)
	
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = models.ValidateStudentData(&student)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Aluno atualizado com sucesso!",
	})
}

func HandleGetStudentByCPF(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	var student models.Student

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, student) 
}
