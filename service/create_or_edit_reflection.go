package service

import (
	"context"
	"fmt"
	"time"

	"github.com/budougumi0617/go_todo_app/auth"
	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
)

type CreateOrEditReflection struct {
	DB   store.Execer
	Repo ReflectionCreator
}

func (a *CreateOrEditReflection) CreateOrEditReflection(
	    ctx context.Context, content string,
			contentType entity.ContentType, dateString string,
			dateType entity.DateType,
		) (*entity.Reflection, error,
	) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	// dateをstringからtime.Time加工する
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, fmt.Errorf("cannot convert dateString to time.Time")
	}

	t := &entity.Reflection{
		UserID:      id,
		Content:     content,
		ContentType: contentType,
		Date:        date,
		DateType:    dateType,
	}
	if err := a.Repo.CreateOrEditReflection(ctx, a.DB, t); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
