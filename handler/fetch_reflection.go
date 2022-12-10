package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/go-playground/validator/v10"
)

type FetchReflection struct {
	Service   FetchReflectionService
	Validator *validator.Validate
}

func (lt *FetchReflection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Date       string            `json:"date" validate:"required"`
		DateType   entity.DateType   `json:"date_type" validate:"required"`
		WeekNumber entity.WeekNumber `json:"week_number"`
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
	reflections, err := lt.Service.FetchReflection(
		ctx, b.Date, b.DateType, b.WeekNumber,
	)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, reflections, http.StatusOK)
}
