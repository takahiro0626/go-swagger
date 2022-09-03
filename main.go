package main

import (
	"net/http"

	_ "go-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Golang Swagger API Document
// @version version1.0
// @description API Document practice
// @termsOfService Document content

// @contact.name API supporter
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name license
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.New()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/users", searchUsers)
	r.GET("/users/:id", findUserById)
	r.POST("/users", createUser)

	r.Run()
}

type User struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"test"`
	Age  int    `json:"age" example:"20"`
}

type UsersResponse struct {
	Users interface{} `json: "users"`
}

type CreateUserRequest struct {
	Name string `json:"name" example:"test"`
	Age  int    `json:"age" example:"20"`
}

type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// searchUsers
// @Summary return users
// @description return users
// @version 1.0
// @tags users
// @produce json
// @accept application/json
// @param name query string false "name"
// @param age query int false "age"
// @router /users [get]
// @Success 200 {object} UsersResponse{users=[]User}
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
func searchUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": "ok"})
}

// findUserById
// @Summary return user by user_id
// @description return user by user_id
// @version 1.0
// @tags users
// @produce json
// @accept application/json
// @param user_id path int true "user_id"
// @router /users/:user_id [get]
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
func findUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "ok"})
}

// createUser
// @Summary create user
// @description create user
// @version 1.0
// @tags users
// @Accept  json
// @produce json
// @accept application/json
// @param user body CreateUserRequest true "user"
// @router /users [post]
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "ok"})
}