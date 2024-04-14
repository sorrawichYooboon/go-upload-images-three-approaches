package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sorrawichYooboon/gouploadimagesthreeapp/core/health"
	imagemanagement "github.com/sorrawichYooboon/gouploadimagesthreeapp/core/image-management"
)

func main() {
	router := gin.Default()

	healthController := health.NewHealthController()
	imageManagementController := imagemanagement.NewImageManagementController()

	router.GET("/ping", healthController.Ping)

	router.POST("/upload-image/direct", imageManagementController.DirectUploadImage)
	router.POST("/upload-image/base64/:image-name", imageManagementController.Base64UploadImage)
	router.POST("/upload-image/multipart", imageManagementController.MultiPartUploadImage)

	router.Run(":8080")
}
