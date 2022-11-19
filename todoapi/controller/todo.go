package controller

import (
	"net/http"
	"todoapi/domain"
	"todoapi/model"

	"github.com/labstack/echo/v4"
)

func (s *storage) CreateTodo(ctx echo.Context) error {
	req := domain.TodoReq{}

	if err := ctx.Bind(&req); err != nil {
		return err
	}
	tx := s.db.Create(model.Todo{
		Task: req.Task,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return ctx.JSON(http.StatusOK, &domain.Message{
		Message: "success",
	})
}
