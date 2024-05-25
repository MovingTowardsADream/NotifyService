package v1

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/usecase"
	"NotifiService/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type sendMessageRoutes struct {
	w usecase.NotifySend
	l *slog.Logger
}

func newSendMessageRoutes(handler *gin.RouterGroup, w usecase.NotifySend, l *slog.Logger) {
	r := &sendMessageRoutes{w, l}

	h := handler.Group("/messages")
	{
		h.POST("/send", r.sendMessageById)
	}
}

func (r *sendMessageRoutes) sendMessageById(c *gin.Context) {
	var notifyRequest entity.NotificationRequest

	if err := c.BindJSON(&notifyRequest); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := r.w.SendNotifyForUser(c.Request.Context(), notifyRequest)
	if err != nil {
		if errors.Is(err, entity.ErrTimeout) {
			c.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}

		r.l.Error("http - v1 - sendMessageById", logger.Err(err))
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, "ok")
}
