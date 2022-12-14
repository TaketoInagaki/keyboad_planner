package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/go-playground/validator/v10"
)

type CreateOrEditContinuationList struct {
	Service   CreateOrEditContinuationService
	Validator *validator.Validate
}

func (at *CreateOrEditContinuationList) ServeHTTP(
	w http.ResponseWriter, r *http.Request,
) {
	ctx := r.Context()
	var body struct {
		ID               entity.ContinuationID   `json:"id"`
		Content          string                  `json:"content" validate:"required"`
		ContinuationType entity.ContinuationType `json:"content_type" validate:"required"`
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
	t, err := at.Service.CreateOrEditContinuationList(
		ctx, body.ID, body.Content, body.ContinuationType,
	)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.ContinuationID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
