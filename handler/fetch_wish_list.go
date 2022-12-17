package handler

import (
	"net/http"
)

type FetchWishList struct {
	Service FetchWishService
}

func (fc *FetchWishList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	wishes, err := fc.Service.FetchWishList(ctx)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, wishes, http.StatusOK)
}
