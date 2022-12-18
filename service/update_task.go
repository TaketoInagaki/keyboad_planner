package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type UpdateTask struct {
	DB   store.Execer
	Repo TaskUpdater
}

func (a *UpdateTask) UpdateTask(
	ctx context.Context, id entity.TaskID, status entity.ActionStatus,
) (*entity.Task, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	ac := &entity.Task{
		ID:     id,
		UserID: userId,
		Status: status,
	}
	if err := a.Repo.UpdateTask(ctx, a.DB, ac); err != nil {
		return nil, fmt.Errorf("failed to edit: %w", err)
	}

	return ac, nil
}
