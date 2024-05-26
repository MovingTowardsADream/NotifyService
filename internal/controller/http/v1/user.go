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

// @Summary     Changing preferences
// @Description Changes user preferences
// @Tags  	    Preferences
// @Accept      json
// @Produce     json
// @Param       input body entity.RequestPreferences true "user preferences"
// @Success     200 {string} Success
// @Failure     400 "Bad request"
// @Failure     500 "Failed to edit preferences"
// @Failure     504 "Waiting time is up"
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
		Message: "Success",
	})
}
