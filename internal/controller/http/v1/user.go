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

// @Summary     Изменение предпочтений
// @Description Изменяет предпочтения пользователя
// @Tags  	    Preferences
// @Success     200 "success"
// @Failure     500 "Не удалось изменить предпочтения"
// @Failure     504 "Время ожидания вышло"
// @Example     { "userId": "bf82a761ab8c5ed627c136571d33cb55", "preferences": { "email":{ "notifyType": "alert", "approval": true }, "phone":{ "notifyType": "alert", "approval": true } } }
// @Router      /user/preferences [put].
func (r *usersRoutes) UserPreferences(c *gin.Context) {
	var preferences entity.RequestPreferences

	if err := c.BindJSON(&preferences); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := r.w.EditPreferences(c.Request.Context(), preferences)
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

	c.JSON(http.StatusOK, response{
		Message: "success",
	})
}
