package service

import (
	"context"
	"fmt"
	"time"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type CreateTask struct {
	DB   store.Execer
	Repo TaskCreator
}

func (a *CreateTask) CreateOrEditTask(
	ctx context.Context, id entity.TaskID, title string,
	dateString string, dateType entity.TaskDateType,
	weekNumber entity.WeekNumber,
) (*entity.Task, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	date, err := convertToTimeTask(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}
	// storeに渡す値をまとめる
	t := &entity.Task{
		ID:         id,
		UserID:     user_id,
		Title:      title,
		Date:       *date,
		DateType:   dateType,
		WeekNumber: weekNumber,
	}
	// 保存or編集
	if t.ID != 0 {
		if err := a.Repo.EditTask(ctx, a.DB, t); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		err := a.Repo.CreateTask(ctx, a.DB, t)
		if err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return t, nil
}

func convertToTimeTask(
	dateString string, dateType entity.TaskDateType, weekNumber entity.WeekNumber,
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
	default:
		return nil, fmt.Errorf("don't match any dateType")
	}
	return &date, nil
}
