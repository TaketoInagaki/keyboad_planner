package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TaketoInagaki/keyboad_planner/entity"
	"github.com/go-playground/validator/v10"
)

type CreateOrEditReflection struct {
	Service   CreateOrEditReflectionService
	Validator *validator.Validate
}

func (at *CreateOrEditReflection) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		Content     string             `json:"content" validate:"required"`
		ContentType entity.ContentType `json:"content_type" validate:"required"`
		Date        string             `json:"date" validate:"required"`
		DateType    entity.DateType    `json:"date_type" validate:"required"`
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
	t, err := at.Service.CreateOrEditReflection(ctx, body.Content, body.ContentType, body.Date , body.DateType)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.ReflectionID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
