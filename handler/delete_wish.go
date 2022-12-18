package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/go-playground/validator/v10"
)

type DeleteWish struct {
	Service   DeleteWishService
	Validator *validator.Validate
}

func (at *DeleteWish) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		ID entity.WishID `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := at.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t, err := at.Service.DeleteWish(ctx, b.ID)
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
