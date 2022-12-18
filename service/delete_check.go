package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type DeleteCheck struct {
	DB   store.Execer
	Repo CheckDeleter
}

func (a *DeleteCheck) DeleteCheck(
	ctx context.Context, id entity.CheckID,
) (*entity.Check, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// storeに渡す値をまとめる
	t := &entity.Check{
		ID:     id,
		UserID: user_id,
	}
	if err := a.Repo.DeleteCheck(ctx, a.DB, t); err != nil {
		return nil, fmt.Errorf("failed to edit: %w", err)
	}
	return t, nil
}
