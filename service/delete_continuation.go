package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type DeleteContinuation struct {
	DB   store.Execer
	Repo ContinuationDeleter
}

func (a *DeleteContinuation) DeleteContinuation(
	ctx context.Context, id entity.ContinuationID,
) (*entity.Continuation, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// storeに渡す値をまとめる
	t := &entity.Continuation{
		ID:     id,
		UserID: user_id,
	}
	if err := a.Repo.DeleteContinuation(ctx, a.DB, t); err != nil {
		return nil, fmt.Errorf("failed to edit: %w", err)
	}
	return t, nil
}
