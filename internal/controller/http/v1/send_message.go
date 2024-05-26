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

// @Summary     Отправка сообщения
// @Description Отправляет сообщение пользователю с заданными полями
// @Tags  	    Notify
// @Success     200 "success"
// @Failure     500 "Не удалось создать нотификацию"
// @Failure     504 "Время ожидания вышло"
// @Example     { "userId": "bf82a761ab8c5ed627c136571d33cb55", "notifyType": "alert", "channels":{ "email":{ "subject": "New alert!", "body": "<html>..." }, "phone":{ "subject": "New alert!", "body": "<html>..." } } }
// @Router      /messages/send [post].
func (r *sendMessageRoutes) sendMessageById(c *gin.Context) {
	var notifyRequest entity.RequestNotification

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

	type response struct {
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, response{
		Message: "success",
	})
}
