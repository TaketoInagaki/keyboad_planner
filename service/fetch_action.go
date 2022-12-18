package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type FetchAction struct {
	DB   store.Queryer
	Repo ActionFetcher
}

type Action struct {
	ID         entity.ActionID   `json:"id"`
	Content    string            `json:"content"`
	Date       string            `json:"date"`
	DateType   entity.DateType   `json:"date_type"`
	WeekNumber entity.WeekNumber `json:"week_number"`
}

type Actions []Action

func (f *FetchAction) FetchAction(
	ctx context.Context, dateString string,
	dateType entity.DateType, weekNumber entity.WeekNumber,
) (Actions, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// 日付をtime.Timeに変換する
	date, err := convertToTimeReflection(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}
	r := &entity.Action{
		UserID:     user_id,
		Date:       *date,
		WeekNumber: weekNumber,
	}
	// storeを使ってデータを取得する
	rs, err := f.Repo.FetchAction(ctx, f.DB, r)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	checks := []Action{}
	for _, r := range rs {
		dateString := convertToStringReflection(r.Date, r.DateType)
		checks = append(checks, Action{
			ID:         r.ID,
			Content:    r.Content,
			Date:       *dateString,
			DateType:   r.DateType,
			WeekNumber: r.WeekNumber,
		})
	}

	return checks, nil
}
