package v1

import (
	"NotifiService/internal/entity"
	"NotifiService/internal/usecase"
	"NotifiService/pkg/logger"
	"context"
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
		h.POST("/send", r.sendAllMessage)
	}
}

func (r *sendMessageRoutes) sendAllMessage(c *gin.Context) {
	err := r.w.SendNotifyForUser(context.Background(), "")
	if err != nil {
		if errors.Is(err, entity.ErrTimeout) {
			c.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}

		r.l.Error("http - v1 - createNewWallet", logger.Err(err))
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, "ok")
}
