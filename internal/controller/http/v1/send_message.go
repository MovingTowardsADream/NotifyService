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

// @Summary     Sending a message
// @Description Sends a message to the user with the specified fields.
// @Tags  	    Notify
// @Accept      json
// @Produce     json
// @Param       input body entity.RequestNotification true "send mess"
// @Success     200 {string} Success
// @Failure     400 "Bad request"
// @Failure     500 "Failed to create notification"
// @Failure     504 "Waiting time is up"
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
		Message: "Success",
	})
}
