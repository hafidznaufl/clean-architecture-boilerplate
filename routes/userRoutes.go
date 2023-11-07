package routes

import (
	"os"
	"peekabook/controller"
	"peekabook/repository"
	"peekabook/service"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	usersGroup := e.Group("users")

	usersGroup.POST("/register", userController.RegisterUserController)
	usersGroup.POST("/login", userController.LoginUserController)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.GET("/:id", userController.GetUserController)
	usersGroup.GET("", userController.GetUsersController)
	usersGroup.GET("", userController.GetUserByNameController)
	usersGroup.PUT("/:id", userController.UpdateUserController)
	usersGroup.PUT("/reset-password", userController.ResetPassword)
	usersGroup.DELETE("/:id", userController.DeleteUserController)

}
