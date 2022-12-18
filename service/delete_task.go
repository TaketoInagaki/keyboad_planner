package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type DeleteTask struct {
	DB   store.Execer
	Repo TaskDeleter
}

func (a *DeleteTask) DeleteTask(
	ctx context.Context, id entity.TaskID,
) (*entity.Task, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// storeに渡す値をまとめる
	t := &entity.Task{
		ID:     id,
		UserID: user_id,
	}
	if err := a.Repo.DeleteTask(ctx, a.DB, t); err != nil {
		return nil, fmt.Errorf("failed to edit: %w", err)
	}
	return t, nil
}
