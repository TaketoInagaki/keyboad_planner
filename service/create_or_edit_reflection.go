package service

import (
	"context"
	"fmt"
	"time"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type CreateOrEditReflection struct {
	DB   store.Execer
	Repo ReflectionCreator
}

func (a *CreateOrEditReflection) CreateOrEditReflection(ctx context.Context, id entity.ReflectionID, content string, contentType entity.ContentType, dateString string, dateType entity.DateType) (*entity.Reflection, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	// dateをstringからtime.Time加工する
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, fmt.Errorf("cannot convert dateString to time.Time")
	}

	t := &entity.Reflection{
		ID:          id,
		UserID:      userId,
		Content:     content,
		ContentType: contentType,
		Date:        date,
		DateType:    dateType,
	}
	if t.ID != 0 {
		if err := a.Repo.EditReflection(ctx, a.DB, t); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		if err := a.Repo.CreateReflection(ctx, a.DB, t); err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return t, nil
}
