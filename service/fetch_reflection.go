package service

import (
	"context"
	"fmt"
	"time"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type FetchReflection struct {
	DB   store.Queryer
	Repo ReflectionFetcher
}

type Reflection struct {
	ID          entity.ReflectionID `json:"id"`
	Content     string              `json:"content"`
	ContentType entity.ContentType  `json:"content_type"`
	Date        string              `json:"date"`
	DateType    entity.DateType     `json:"date_type"`
	WeekNumber  entity.WeekNumber   `json:"week_number"`
}

type Reflections []Reflection

func (f *FetchReflection) FetchReflection(
	ctx context.Context, dateString string,
	dateType entity.DateType, weekNumber entity.WeekNumber,
) (Reflections, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// 日付をtime.Timeに変換する
	date, err := convertToTimeReflection(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}
	r := &entity.Reflection{
		UserID:     user_id,
		Date:       *date,
		WeekNumber: weekNumber,
	}
	// storeを使ってデータを取得する
	rs, err := f.Repo.FetchReflection(ctx, f.DB, r)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	reflections := []Reflection{}
	for _, r := range rs {
		dateString := convertToStringReflection(r.Date, r.DateType)
		reflections = append(reflections, Reflection{
			ID:          r.ID,
			Content:     r.Content,
			ContentType: r.ContentType,
			Date:        *dateString,
			DateType:    r.DateType,
			WeekNumber:  r.WeekNumber,
		})
	}

	return reflections, nil
}

func convertToStringReflection(date time.Time, dateType entity.DateType) *string {
	// time.Timeの日付を文字列に変換する
	var dateString string
	switch dateType {
	case "Monthly":
		dateString = date.Format(time.RFC3339Nano)[0:7]
	case "Yearly", "Weekly":
		dateString = date.Format(time.RFC3339Nano)[0:4]
	case "Daily":
		dateString = date.Format(time.RFC3339Nano)[0:10]
	}
	return &dateString
}
