package api

import (
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/", index)

	routes := r.Group("/api/v1")
	routes.Use(authorization())
	routes.POST("/password", addPassword)
	routes.GET("/password", passwordList)

	routes.GET("/password/:password_id", passwordDetails)
	routes.PUT("/password/:password_id", updatePassword)
	routes.DELETE("/password/:password_id", deletePassword)

	return r
}
