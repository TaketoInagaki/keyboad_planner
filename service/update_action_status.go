package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type UpdateAction struct {
	DB   store.Execer
	Repo ActionUpdater
}

func (a *UpdateAction) UpdateAction(
	ctx context.Context, id entity.ActionID, status entity.ActionStatus,
) (*entity.Action, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	ac := &entity.Action{
		ID:     id,
		UserID: userId,
		Status: status,
	}
	if err := a.Repo.UpdateAction(ctx, a.DB, ac); err != nil {
		return nil, fmt.Errorf("failed to edit: %w", err)
	}

	return ac, nil
}
