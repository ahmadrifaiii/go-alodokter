package order

import (
	"go-alodokter/model"
	"go-alodokter/module/v1/order/usecase"
	"go-alodokter/utl/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.orderList).Name = "order-list"
	group.GET("/detail/:param", m.orderDetail).Name = "order-detail"
	group.POST("/register", m.createOrder).Name = "order-register"
	group.PUT("/update/:param", m.updateOrder).Name = "order-update"
	group.DELETE("/delete/:param", m.deleteOrder).Name = "order-delete"
}

// @Summary Get List Order
// @Description get the status of server.
// @Tags Order Management System - Order
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/order/list [get]
func (m *Module) orderList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	resp, err := usecase.OrderList(m.Config)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}

// @Summary Create Order
// @Description get the status of server.
// @Tags Order Management System - Order
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/order/create [post]
func (m *Module) createOrder(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		email     = c.Get("email").(string)
		usr       = model.User{}
	)

	err := c.Bind(&usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	usr.CreatedBy = &email
	usr.UpdatedBy = &email

	resp, err := usecase.UserRegister(m.Config, &usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "user has been registered",
		Data:    resp,
	})
}

// @Summary Update User
// @Description get the status of server.
// @Tags Order Management System - Order
// @Accept */*
// @Produce json
// @Param param path int true "User Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/order/update/{param} [put]
func (m *Module) updateOrder(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		usr       = model.User{}
	)

	id := c.Param("param")
	usr.Id = id

	err := c.Bind(&usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	resp, err := usecase.UserUpdate(m.Config, &usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    resp,
	})
}

// @Summary Delete the User
// @Description get the status of server.
// @Tags Order Management System - Order
// @Accept */*
// @Produce json
// @Param param path string true "order id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/order/delete/{param} [delete]
func (m *Module) deleteOrder(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		usr       = model.User{}
	)

	id := c.Param("param")
	usr.Id = id

	err := c.Bind(&usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	_, err = usecase.UserDelete(m.Config, &usr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    nil,
	})
}

// @Summary Detail of User
// @Description get the status of server.
// @Tags Order Management System - Order
// @Accept */*
// @Produce json
// @Param param path int true "User Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/order/detail/{param} [get]
func (m *Module) orderDetail(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)
	id, _ := strconv.Atoi(c.Param("param"))
	resp, err := usecase.UserDetail(m.Config, id)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
		Error:   err,
	})
}
