package auth

import (
	"go-alodokter/model"
	"go-alodokter/module/v1/auth/usecase"
	"go-alodokter/utl/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.POST("/login", m.login)
}

// @Summary User Administrator Login
// @Description get the status of server.
// @Tags Admin - Login
// @Accept json
// @Produce json
// @Param body body model.LoginPayload true "model.LoginPayload"
// @Success 202 {object} model.Response{}
// @Failuer 400 {object} model.Response{}
// @Router /api/v1/auth/login [post]
func (m *Module) login(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		payload   = model.LoginPayload{}
	)

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Data:    nil,
			Error:   err,
		})
	}

	resp, err := usecase.Login(m.Config, &payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Data:    nil,
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusAccepted,
		Message: "login success",
		Data:    resp,
	})
}
