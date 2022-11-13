package handler

import (
	"net/http"
	"encoding/json"

	"github.com/TaketoInagaki/keyboard_planner/entity"
    "github.com/go-playground/validator/v10"
)

type ListTask struct {
	Service ListTasksService
	Validator *validator.Validate
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Date       string              `json:"date" validate:"required"`
		DateType   entity.TaskDateType `json:"date_type" validate:"required"`
		WeekNumber entity.WeekNumber   `json:"week_number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := lt.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	tasks, err := lt.Service.ListTasks(
		ctx, b.Date, b.DateType, b.WeekNumber,
	)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, tasks, http.StatusOK)
}
