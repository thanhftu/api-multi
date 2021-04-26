package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thanhftu/api-multi/controller"
)

var (
	router = gin.Default()
)

func main() {
	// router.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowAllOrigins:  false,
	// 	AllowOriginFunc:  func(origin string) bool { return true },
	// 	MaxAge:           86400,
	// }))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	router.Use(cors.New(config))
	router.GET("/", controller.GetFibFromDB)
	router.POST("/api/values", controller.GetWorkerResultHandler)
	router.GET("/api/values/latest", controller.GetLatestFibHandler)
	router.DELETE("/api/values/:id", controller.DeleteFibHandler)
	router.GET("/allfib", controller.GetAllFinController)
	router.Run(":8081")
}
