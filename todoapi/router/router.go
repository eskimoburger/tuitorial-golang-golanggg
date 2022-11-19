package router

import (
	"net/http"
	"todoapi/domain"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {

	e.GET("/health", func(c echo.Context) error {

		return c.JSON(http.StatusOK, domain.Health{
			Message: "ok",
		})
	})

	e.POST("/todo", func(c echo.Context) error {
		return nil
	})
	e.GET("/todo/:id", func(c echo.Context) error { return nil })
	e.GET("/todos", func(c echo.Context) error { return nil })
	e.PUT("/todo/:id", func(c echo.Context) error { return nil })
	e.DELETE("/todo/:id", func(c echo.Context) error { return nil })

}
