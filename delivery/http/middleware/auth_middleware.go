package middleware

import (
	"EmployeeManagementApp/config"
	delivery_helper "EmployeeManagementApp/delivery/util"
	"EmployeeManagementApp/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	requestID, _ := c.Get("request_id")

	var tokenStr string
	cookies := c.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "token" {
			tokenStr = cookie.Value
		}
	}

	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, delivery_helper.WriteErrorResponse(requestID.(string), "Token is missing"))
		c.Abort()
		return
	}

	token, errs := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetServer().SignatureKey), nil
	})

	if errs != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, delivery_helper.WriteErrorResponse(requestID.(string), "Unauthorized"))
		c.Abort()
		return
	}

	defer func() {
		if errs != nil {
			util.LogError(util.DefaultGenerateLogModel(500, errs.Error()).LoggerZapFieldObject())
		}
	}()

	claims, _ := token.Claims.(jwt.MapClaims)
	c.Set("claims", claims)

	c.Next()
}
