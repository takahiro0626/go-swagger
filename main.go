package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-swagger/docs"
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

	r.GET("/test", test)
	r.Run()
}

// @description test api detail
// @version 1.0
// @accept application/x-json-stream
// @param none query string false "not require"
// @router /test/ [get]
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
