package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type CreateOrEditContinuationList struct {
	DB   store.Execer
	Repo ContinuationCreator
}

func (a *CreateOrEditContinuationList) CreateOrEditContinuationList(
	ctx context.Context, id entity.ContinuationID, content string,
	contentType entity.ContinuationType,
) (*entity.Continuation, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	con := &entity.Continuation{
		ID:             id,
		UserID:         userId,
		Content:        content,
		ContinuationType: contentType,
	}
	if con.ID != 0 {
		if err := a.Repo.EditContinuation(ctx, a.DB, con); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		if err := a.Repo.CreateContinuation(ctx, a.DB, con); err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return con, nil
}
