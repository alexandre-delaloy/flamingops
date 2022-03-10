package controllers

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/middlewares"
	// "github.com/blyndusk/flamingops/pkg/models"
)

func CreateSecret(c *gin.Context) {
	middlewares.CreateSecret(c)
}

func UpdateSecret(c *gin.Context) {
	middlewares.UpdateSecret(c)
}

func DeleteSecret(c *gin.Context) {
	middlewares.DeleteSecret(c)
}
