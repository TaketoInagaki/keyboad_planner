package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type CreateOrEditAction struct {
	DB   store.Execer
	Repo ActionCreator
}

func (a *CreateOrEditAction) CreateOrEditAction(
	ctx context.Context, id entity.ActionID, content string, dateString string,
	dateType entity.DateType, weekNumber entity.WeekNumber,
) (*entity.Action, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	// 日付をtime.Timeに変換する
	date, err := convertToTimeReflection(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}

	ref := &entity.Action{
		ID:         id,
		UserID:     userId,
		Content:    content,
		Date:       *date,
		DateType:   dateType,
		WeekNumber: weekNumber,
	}
	if ref.ID != 0 {
		if err := a.Repo.EditAction(ctx, a.DB, ref); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		if err := a.Repo.CreateAction(ctx, a.DB, ref); err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return ref, nil
}
