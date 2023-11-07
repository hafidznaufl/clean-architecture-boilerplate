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

func AdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository, validate)
	adminController := controller.NewAdminController(adminService)

	adminsGroup := e.Group("admin")

	adminsGroup.POST("", adminController.RegisterAdminController)
	adminsGroup.POST("/login", adminController.LoginAdminController)

	adminsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	adminsGroup.GET("/:id", adminController.GetAdminController)
	adminsGroup.GET("", adminController.GetAdminsController)
	adminsGroup.GET("", adminController.GetAdminByNameController)
	adminsGroup.PUT("/:id", adminController.UpdateAdminController)
	adminsGroup.DELETE("/:id", adminController.DeleteAdminController)
}
