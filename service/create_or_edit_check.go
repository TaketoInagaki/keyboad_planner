package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type CreateOrEditCheck struct {
	DB    store.Execer
	Repo  CheckCreator
}

func (a *CreateOrEditCheck) CreateOrEditCheck(
	ctx context.Context, id entity.CheckID, content string, dateString string,
	dateType entity.DateType, weekNumber entity.WeekNumber,
) (*entity.Check, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	// 日付をtime.Timeに変換する
	date, err := convertToTimeReflection(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}

	ref := &entity.Check{
		ID:         id,
		UserID:     userId,
		Content:    content,
		Date:       *date,
		DateType:   dateType,
		WeekNumber: weekNumber,
	}
	if ref.ID != 0 {
		if err := a.Repo.EditCheck(ctx, a.DB, ref); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		if err := a.Repo.CreateCheck(ctx, a.DB, ref); err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return ref, nil
}
