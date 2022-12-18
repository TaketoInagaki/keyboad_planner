package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/go-playground/validator/v10"
)

type CreateOrEditAction struct {
	Service   CreateOrEditActionService
	Validator *validator.Validate
}

func (at *CreateOrEditAction) ServeHTTP(
	w http.ResponseWriter, r *http.Request,
) {
	ctx := r.Context()
	var body struct {
		ID         entity.ActionID   `json:"id"`
		Content    string            `json:"content" validate:"required"`
		Date       string            `json:"date" validate:"required"`
		DateType   entity.DateType   `json:"date_type" validate:"required"`
		WeekNumber entity.WeekNumber `json:"week_number"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := at.Validator.Struct(body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	t, err := at.Service.CreateOrEditAction(
		ctx, body.ID, body.Content,
		body.Date, body.DateType, body.WeekNumber,
	)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.ActionID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
