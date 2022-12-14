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

func (a *CreateOrEditReflection) CreateOrEditReflection(
	ctx context.Context, id entity.ReflectionID, content string,
	contentType entity.ReflectionType, dateString string,
	dateType entity.DateType, weekNumber entity.WeekNumber,
) (*entity.Reflection, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	// 日付をtime.Timeに変換する
	date, err := convertToTimeReflection(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}

	ref := &entity.Reflection{
		ID:             id,
		UserID:         userId,
		Content:        content,
		ReflectionType: contentType,
		Date:           *date,
		DateType:       dateType,
		WeekNumber:     weekNumber,
	}
	if ref.ID != 0 {
		if err := a.Repo.EditReflection(ctx, a.DB, ref); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		if err := a.Repo.CreateReflection(ctx, a.DB, ref); err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return ref, nil
}

func convertToTimeReflection(
	dateString string, dateType entity.DateType, weekNumber entity.WeekNumber,
) (*time.Time, error) {
	// 日付をtime.Timeに変換する
	var date time.Time
	var err error
	switch dateType {
	case "Monthly":
		date, err = time.Parse("2006-01", dateString[0:7])
		if err != nil {
			return nil, fmt.Errorf("cannot convert dateString to time.Time")
		}
	case "Yearly", "Weekly":
		date, err = time.Parse("2006", dateString[0:4])
		if err != nil {
			return nil, fmt.Errorf("cannot convert dateString to time.Time")
		}
		if dateType == "Weekly" && weekNumber == 0 {
			return nil, fmt.Errorf("when weekly week_number is required")
		}
	case "Daily":
		date, err = time.Parse("2006-01-02", dateString)
		if err != nil {
			return nil, fmt.Errorf("cannot convert dateString to time.Time")
		}
	default:
		return nil, fmt.Errorf("doesn't match any dateType")
	}
	return &date, nil
}
