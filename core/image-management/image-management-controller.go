package imagemanagement

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type IImageManagementController interface {
	DirectUploadImage(c *gin.Context)
	Base64UploadImage(c *gin.Context)
	MultiPartUploadImage(c *gin.Context)
}

type ImageManagementController struct {
}

func NewImageManagementController() *ImageManagementController {
	return &ImageManagementController{}
}

func (controller *ImageManagementController) DirectUploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	out, err := os.Create("assets/images/" + header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (controller *ImageManagementController) Base64UploadImage(c *gin.Context) {
	imageData, _ := c.GetPostForm("image")
	if imageData == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing image data"})
		return
	}

	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(imageData)))

	imgBuffer := bytes.NewBuffer(nil)
	_, err := io.Copy(imgBuffer, reader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error decoding Base64 image"})
		return
	}

	img, err := jpeg.Decode(imgBuffer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error decoding image"})
		return
	}

	fileName := "assets/images/" + c.Param("image-name") + ".jpg"

	file, err := os.Create(fileName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error creating image file"})
		return
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error saving image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}

func (controller *ImageManagementController) MultiPartUploadImage(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["images"]
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer src.Close()

		out, err := os.Create("assets/images/" + file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, src)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})
}
