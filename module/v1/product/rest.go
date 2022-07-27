package product

import (
	"go-alodokter/model"
	"go-alodokter/module/v1/product/usecase"

	"go-alodokter/utl/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.productList).Name = "product-list"
	group.GET("/detail/:param", m.userDetail).Name = "product-detail"
	group.POST("/create", m.createProduct).Name = "product-create"
	group.PUT("/update/:param", m.userUpdate).Name = "product-update"
	group.DELETE("/delete/:param", m.productDelete).Name = "product-delete"
}

// @Summary Get List Product
// @Description get the status of server.
// @Tags Product Management - Product
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/product/list [get]
func (m *Module) productList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := usecase.ProductList(m.Config)
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

// @Summary Create Product
// @Description get the status of server.
// @Tags Product Management - Product
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/product/create [post]
func (m *Module) createProduct(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		email     = c.Get("email").(string)
		payload   = model.Product{}
	)

	err := c.Bind(&payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	payload.CreatedBy = &email
	payload.UpdatedBy = &email

	resp, err := usecase.CreateNewProduct(m.Config, &payload)
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
		Message: "created product success",
		Data:    resp,
	})
}

// @Summary Update Product
// @Description get the status of server.
// @Tags Product Management - Product
// @Accept */*
// @Produce json
// @Param param path string true "Product Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/product/update/{param} [put]
func (m *Module) userUpdate(c echo.Context) error {
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

// @Summary Delete Product
// @Description get the status of server.
// @Tags Product Management - Product
// @Accept */*
// @Produce json
// @Param param path string true "Product Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/product/delete/{param} [delete]
func (m *Module) productDelete(c echo.Context) error {
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

// @Summary Detail of Produt
// @Description get the status of server.
// @Tags Product Management - Product
// @Accept */*
// @Produce json
// @Param param path string true "Product Id"
// @Success 200 {interface} model.Response{}
// @Router /api/v1/product/detail/{param} [get]
func (m *Module) userDetail(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)
	id, _ := strconv.Atoi(c.Param("param"))
	resp, err := usecase.ProductDetail(m.Config, id)
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
