package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type FetchContinuationList struct {
	DB   store.Queryer
	Repo ContinuationFetcher
}

type Continuation struct {
	UserID           entity.UserID           `json:"user_id"`
	ID               entity.ContinuationID   `json:"id"`
	Content          string                  `json:"content"`
	ContinuationType entity.ContinuationType `json:"content_type"`
	DateType         entity.DateType         `json:"date_type"`
	WeekNumber       entity.WeekNumber       `json:"week_number"`
}

type Continuations struct {
	Continue []Continuation
	Begin    []Continuation
	Quit     []Continuation
}

func (f *FetchContinuationList) FetchContinuationList(
	ctx context.Context,
) (*Continuations, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	c := &entity.Continuation{
		UserID: user_id,
	}
	// storeを使ってデータを取得する
	cs, err := f.Repo.FetchContinuation(ctx, f.DB, c)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	var continuations Continuations
	for _, c := range cs {
		// コンテンツタイプによって入れる場所を変える
		switch c.ContinuationType {
		case "Continue":
			continuations.Continue = append(continuations.Continue, Continuation{
				ID:               c.ID,
				Content:          c.Content,
				ContinuationType: c.ContinuationType,
			})
		case "Begin":
			continuations.Begin = append(continuations.Begin, Continuation{
				ID:               c.ID,
				Content:          c.Content,
				ContinuationType: c.ContinuationType,
			})
		case "Quit":
			continuations.Quit = append(continuations.Quit, Continuation{
				ID:               c.ID,
				Content:          c.Content,
				ContinuationType: c.ContinuationType,
			})
		}
	}

	return &continuations, nil
}
