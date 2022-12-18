package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type FetchCheck struct {
	DB   store.Queryer
	Repo CheckFetcher
}

type Check struct {
	ID         entity.CheckID    `json:"id"`
	Content    string            `json:"content"`
	Date       string            `json:"date"`
	DateType   entity.DateType   `json:"date_type"`
	WeekNumber entity.WeekNumber `json:"week_number"`
}

type Checks []Check

func (f *FetchCheck) FetchCheck(
	ctx context.Context, dateString string,
	dateType entity.DateType, weekNumber entity.WeekNumber,
) (Checks, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// 日付をtime.Timeに変換する
	date, err := convertToTimeReflection(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}
	r := &entity.Check{
		UserID:     user_id,
		Date:       *date,
		WeekNumber: weekNumber,
	}
	// storeを使ってデータを取得する
	rs, err := f.Repo.FetchCheck(ctx, f.DB, r)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	checks := []Check{}
	for _, r := range rs {
		dateString := convertToStringReflection(r.Date, r.DateType)
		checks = append(checks, Check{
			ID:         r.ID,
			Content:    r.Content,
			Date:       *dateString,
			DateType:   r.DateType,
			WeekNumber: r.WeekNumber,
		})
	}

	return checks, nil
}
