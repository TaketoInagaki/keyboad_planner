package handler

import (
	"net/http"
)

type FetchContinuationList struct {
	Service   FetchContinuationService
}

func (fc *FetchContinuationList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	continuations, err := fc.Service.FetchContinuationList(ctx)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, continuations, http.StatusOK)
}
