package handler

import (
	delivery_helper "EmployeeManagementApp/delivery/util"
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	employeeUsecase domain.EmployeeUsecase
}

func NewEmployeeHandler(t domain.EmployeeUsecase) EmployeeHandler {
	return EmployeeHandler{t}
}

func (input EmployeeHandler) StoreEmployeeHandler(c *gin.Context) {
	var (
		fileName  = "employee_handler.go"
		funcName  = "StoreEmployeeHandler"
		employee  domain.EmployeeRequest
		requestID string
		errs      error
		err       util.ErrorModel
	)

	defer delivery_helper.WriteLogStdout(errs, &err, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, delivery_helper.WriteErrorResponse(requestID, "Claims not found"))
		return
	}

	if errs = c.ShouldBindJSON(&employee); errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	var userID int64
	myClaims, ok := claims.(jwt.MapClaims)
	if ok {
		userIDTemp := myClaims["user_id"]
		userID = int64(userIDTemp.(float64))
	}

	err = input.employeeUsecase.Add(c, domain.ContextModel{UserLoginID: userID}, &employee)
	if err.Err != nil {
		c.JSON(err.Code, delivery_helper.WriteErrorResponse(requestID, err.Err.Error()))
		return
	}

	c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Success Store Data!", nil))
}

func (input EmployeeHandler) UpdateEmployeeHandler(c *gin.Context) {
	var (
		fileName  = "employee_handler.go"
		funcName  = "UpdateEmployeeHandler"
		requestID string
		employee  domain.EmployeeUpdateRequest
		errs      error
		err       util.ErrorModel
	)

	defer delivery_helper.WriteLogStdout(errs, &err, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, delivery_helper.WriteErrorResponse(requestID, "Claims not found"))
		return
	}

	idParam := c.Param("id")
	id, errs := strconv.Atoi(idParam)
	if errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	if errs = c.ShouldBindJSON(&employee); errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	var (
		userID int64
	)

	myClaims, ok := claims.(jwt.MapClaims)
	if ok {
		userIDTemp := myClaims["user_id"]
		userID = int64(userIDTemp.(float64))
	}

	employee.ID = int64(id)
	err = input.employeeUsecase.Update(c, domain.ContextModel{
		UserLoginID: userID,
	}, &employee)
	if err.Err != nil {
		c.JSON(err.Code, delivery_helper.WriteErrorResponse(requestID, err.Err.Error()))
		return
	}

	c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Success Update!", nil))
}

func (input EmployeeHandler) DeleteEmployeeHandler(c *gin.Context) {
	var (
		fileName  = "employee_handler.go"
		funcName  = "DeleteEmployeeHandler"
		requestID string
		errs      error
		err       util.ErrorModel
	)

	defer delivery_helper.WriteLogStdout(errs, &err, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, delivery_helper.WriteErrorResponse(requestID, "Claims not found"))
		return
	}

	idParam := c.Param("id")
	id, errs := strconv.Atoi(idParam)
	if errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	var (
		userID int64
	)

	myClaims, ok := claims.(jwt.MapClaims)
	if ok {
		userIDTemp := myClaims["user_id"]
		userID = int64(userIDTemp.(float64))
	}

	err = input.employeeUsecase.Delete(c, domain.ContextModel{
		UserLoginID: userID,
	}, &domain.EmployeeRequest{ID: int64(id)})
	if err.Err != nil {
		c.JSON(err.Code, delivery_helper.WriteErrorResponse(requestID, err.Err.Error()))
		return
	}

	c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Success Delete!", nil))
}

func (input EmployeeHandler) FetchEmployeeHandler(c *gin.Context) {
	var (
		fileName   = "employee_handler.go"
		funcName   = "FetchEmployeeHandler"
		employee   []domain.ListEmployeeResponse
		pagination domain.Pagination
		requestID  string
		errs       error
		err        util.ErrorModel
	)

	defer delivery_helper.WriteLogStdout(errs, &err, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	_, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, delivery_helper.WriteErrorResponse(requestID, "Claims not found"))
		return
	}

	if errs = c.ShouldBindQuery(&pagination); errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	employee, err = input.employeeUsecase.Fetch(c, &domain.EmployeeRequest{GetListParameter: domain.GetListParameter{
		Page:   int64(pagination.Page),
		Limit:  int64(pagination.Limit),
		Filter: pagination.Filter,
		Order:  pagination.Order,
	}})
	if err.Err != nil {
		c.JSON(err.Code, delivery_helper.WriteErrorResponse(requestID, err.Err.Error()))
		return
	}

	c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Success Fetch Data!", employee))
}

func (input EmployeeHandler) ViewEmployeeHandler(c *gin.Context) {
	var (
		fileName  = "employee_handler.go"
		funcName  = "ViewEmployeeHandler"
		employee  domain.ViewEmployeeResponse
		requestID string
		errs      error
		err       util.ErrorModel
	)

	defer delivery_helper.WriteLogStdout(errs, &err, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	_, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, delivery_helper.WriteErrorResponse(requestID, "Claims not found"))
		return
	}

	idParam := c.Param("id")
	id, errs := strconv.Atoi(idParam)
	if errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	employee, err = input.employeeUsecase.GetByID(c, int64(id))
	if err.Err != nil {
		c.JSON(err.Code, delivery_helper.WriteErrorResponse(requestID, err.Err.Error()))
		return
	}

	c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Success Get Detail!", employee))
}
