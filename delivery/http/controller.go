package controller

import (
	"EmployeeManagementApp/app/serverconfig"
	"EmployeeManagementApp/config"
	"EmployeeManagementApp/delivery/http/handler"
	"EmployeeManagementApp/delivery/http/middleware"
	employee_repository "EmployeeManagementApp/repository/employee_repository/pgsql"
	user_repository "EmployeeManagementApp/repository/user_repository/pgsql"
	"EmployeeManagementApp/usecase/employee_usecase"
	"EmployeeManagementApp/usecase/user_usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Controller() {
	prefixPath := config.ApplicationConfiguration.GetServer().PrefixPath
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.BasicMiddleware)

	// For set init service
	userHandler := handler.NewUsersHandler(user_usecase.NewUsersUsecase(user_repository.NewPgsqlUserRepository(serverconfig.ServerAttribute.DBConnection), 10*time.Second))
	employeeHandler := handler.NewEmployeeHandler(employee_usecase.NewEmployeeUsecase(employee_repository.NewPgsqlEmployeeRepository(serverconfig.ServerAttribute.DBConnection), 10*time.Second))

	// WhiteList API no authorization here (public)
	whiteListAPI := r.Group(prefixPath + "/oauth")
	whiteListAPI.POST("/login", userHandler.LoginHandler)
	whiteListAPI.POST("/register", userHandler.RegisterHandler)
	whiteListAPI.POST("/logout", userHandler.LogoutHandler)

	// Private API Employee authorization needed (private)
	privateAPIEmployee := r.Group(prefixPath + "/employee")
	privateAPIEmployee.Use(middleware.AuthMiddleware)
	privateAPIEmployee.POST("", employeeHandler.StoreEmployeeHandler)
	privateAPIEmployee.GET("", employeeHandler.FetchEmployeeHandler)
	privateAPIEmployee.GET("/:id", employeeHandler.ViewEmployeeHandler)
	privateAPIEmployee.PUT("/:id", employeeHandler.UpdateEmployeeHandler)
	privateAPIEmployee.DELETE("/:id", employeeHandler.DeleteEmployeeHandler)

	_ = r.Run(fmt.Sprintf(`:%s`, config.ApplicationConfiguration.GetServer().Port))
}
