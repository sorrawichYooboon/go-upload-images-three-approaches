package health

import "github.com/gin-gonic/gin"

type IHealthController interface {
	Ping(c *gin.Context)
}

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
