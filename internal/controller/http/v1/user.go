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

type usersRoutes struct {
	w usecase.EditInfo
	l *slog.Logger
}

func newUsersRoutes(handler *gin.RouterGroup, w usecase.EditInfo, l *slog.Logger) {
	r := &usersRoutes{w, l}

	h := handler.Group("/user")
	{
		h.PUT("/preferences", r.UserPreferences)
	}
}

func (r *usersRoutes) UserPreferences(c *gin.Context) {
	var preferences entity.UserPreferences

	if err := c.BindJSON(&preferences); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := r.w.EditPreferences(context.Background(), preferences)
	if err != nil {
		if errors.Is(err, entity.ErrTimeout) {
			c.AbortWithStatus(http.StatusGatewayTimeout)
			return
		}

		r.l.Error("http - v1 - UserPreferences", logger.Err(err))
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	type response struct {
		Message string `json:"message"`
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}
