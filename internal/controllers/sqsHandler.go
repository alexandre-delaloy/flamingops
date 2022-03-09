package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/flamingops/internal/queue"
	"github.com/blyndusk/flamingops/pkg/models"
)

func ManuallySendMessage(c *gin.Context) {
	queue.HandleMessageCreation(c, &input)
}
