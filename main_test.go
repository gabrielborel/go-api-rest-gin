package main

import (
    "bytes"
    "encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"api-rest-gin/controllers"
	"api-rest-gin/database"
	"api-rest-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

var ID int

func CreateStudentMock() {
	student := models.Student{Name: "Gabriel", CPF: "19252283714", RG: "123456789"}
	database.DB.Create(&student)

	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetAllStudentsHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	routes := SetupTestRoutes()
	routes.GET("/students", controllers.HandleGetAllStudents)

	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	routes.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByCpfHandler(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	routes := SetupTestRoutes()
	routes.GET("/students/cpf/:cpf", controllers.HandleGetStudentByCPF)

	request, _ := http.NewRequest("GET", "/students/cpf/19252283714", nil)
	response := httptest.NewRecorder()

	routes.ServeHTTP(response, request)

	var mockedStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockedStudent)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "19252283714", mockedStudent.CPF)
}

func TestGetStudentById(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	routes := SetupTestRoutes()
	routes.GET("/students/:id", controllers.HandleGetStudentById)

	path := "/students/" + strconv.Itoa(ID)

	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	routes.ServeHTTP(response, request)

	var mockedStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &mockedStudent)

	assert.Equal(t, "Gabriel", mockedStudent.Name)
	assert.Equal(t, "19252283714", mockedStudent.CPF)
	assert.Equal(t, "123456789", mockedStudent.RG)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectWithDatabase()
	CreateStudentMock()

	routes := SetupTestRoutes()
	routes.DELETE("/students/:id", controllers.HandleDeleteStudent)
	path := "/students/" + strconv.Itoa(ID)

	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()

	routes.ServeHTTP(response, request)

	fmt.Println(response.Code)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateStudent(t *testing.T) {
    database.ConnectWithDatabase()
    CreateStudentMock()
    defer DeleteStudentMock()

    student := models.Student{Name: "Gabriel", CPF: "12312312312", RG: "321321321"}
    jsonStudent, _ := json.Marshal(student)

    routes := SetupTestRoutes()
    routes.PATCH("/students/:id", controllers.HandleUpdateStudent)
    path := "/students/" + strconv.Itoa(ID)

    request, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonStudent))
    response := httptest.NewRecorder()
    routes.ServeHTTP(response, request)

    var updatedStudent models.Student
    json.Unmarshal(response.Body.Bytes(), &updatedStudent)


    assert.Equal(t, student.CPF, updatedStudent.CPF)
    assert.Equal(t, student.RG, updatedStudent.RG)
}
