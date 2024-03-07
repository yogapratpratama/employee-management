package handler

import (
	"EmployeeManagementApp/config"
	delivery_helper "EmployeeManagementApp/delivery/util"
	"EmployeeManagementApp/domain"
	"EmployeeManagementApp/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UsersHandler struct {
	userUsecase domain.UserUsecase
}

func NewUsersHandler(u domain.UserUsecase) UsersHandler {
	return UsersHandler{u}
}

func (input UsersHandler) LogoutHandler(c *gin.Context) {
	cookies := c.Request.Cookies()

	var (
		fileName  = "user_handler.go"
		funcName  = "LogoutHandler"
		jwtToken  *http.Cookie
		requestID string
	)

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)

	defer delivery_helper.WriteLogStdout(nil, &util.ErrorModel{}, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	for _, itemCookie := range cookies {
		if itemCookie.Name == "token" {
			jwtToken = itemCookie
			break
		}
	}

	if jwtToken != nil {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "token",
			Value:    "",
			HttpOnly: true,
			Path:     "/",
			MaxAge:   int(time.Now().Add(-1 * time.Hour).Unix()),
		})
		c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Logout Success!!", nil))
	} else {
		c.JSON(http.StatusNotFound, delivery_helper.WriteErrorResponse(requestID, "Cookies Not Found/Unauthorized"))
	}
}

func (input UsersHandler) LoginHandler(c *gin.Context) {
	var (
		fileName  = "user_handler.go"
		funcName  = "LoginHandler"
		login     domain.Login
		errs      error
		err       util.ErrorModel
		requestID string
	)

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	if errs = c.ShouldBindJSON(&login); errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	defer delivery_helper.WriteLogStdout(errs, &util.ErrorModel{}, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	m := make(map[string]string)
	m[login.Username] = login.Password

	res, err := input.userUsecase.GetLogin(c, &domain.User{Username: login.Username})
	if err.Err != nil {
		return
	}

	if util.CheckPassword(login.Password, res.Password) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = res.ID
		claims["username"] = login.Username
		claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

		var (
			tokenStr  string
			signature = config.ApplicationConfiguration.GetServer().SignatureKey
		)

		tokenStr, errs = token.SignedString([]byte(signature))
		if errs != nil {
			c.JSON(http.StatusInternalServerError, delivery_helper.WriteErrorResponse(requestID, "Failed to create token!"))
			return
		}

		delivery_helper.SetJWTTokenCookie(c, tokenStr, claims)
		c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Token has generated!", nil))
	} else {
		c.JSON(http.StatusUnauthorized, delivery_helper.WriteErrorResponse(requestID, "Invalid credential!"))
	}
}

func (input UsersHandler) RegisterHandler(c *gin.Context) {
	var (
		fileName  = "user_handler.go"
		funcName  = "RegisterHandler"
		user      domain.User
		requestID string
		errs      error
		err       util.ErrorModel
	)

	defer delivery_helper.WriteLogStdout(errs, &err, requestID, fmt.Sprintf(`[%s, %s]`, fileName, funcName))

	requestIDTemp, _ := c.Get("request_id")
	requestID = requestIDTemp.(string)
	if errs = c.ShouldBindJSON(&user); errs != nil {
		c.JSON(http.StatusBadRequest, delivery_helper.WriteErrorResponse(requestID, errs.Error()))
		return
	}

	err = input.userUsecase.Store(c, &user)
	if err.Err != nil {
		c.JSON(err.Code, delivery_helper.WriteErrorResponse(requestID, err.Err.Error()))
		return
	}

	c.JSON(http.StatusOK, delivery_helper.WriteSuccessResponse(requestID, "Success Register!", nil))
}
