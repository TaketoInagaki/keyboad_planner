package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/go-playground/validator/v10"
)

type CreateOrEditWishList struct {
	Service   CreateOrEditWishService
	Validator *validator.Validate
}

func (cw *CreateOrEditWishList) ServeHTTP(
	w http.ResponseWriter, r *http.Request,
) {
	ctx := r.Context()
	var body struct {
		ID      entity.WishID `json:"id"`
		Content string        `json:"content" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := cw.Validator.Struct(body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	t, err := cw.Service.CreateOrEditWishList(
		ctx, body.ID, body.Content,
	)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.WishID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
