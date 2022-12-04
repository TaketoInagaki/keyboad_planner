package service

import (
	"context"
	"fmt"
	"time"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

type Task struct {
	ID         entity.TaskID       `json:"id"`
	Title      string              `json:"title"`
	Date       string              `json:"date"`
	DateType   entity.TaskDateType `json:"date_type"`
	WeekNumber entity.WeekNumber   `json:"week_number"`
}

type Tasks []Task

func (l *ListTask) ListTasks(
	ctx context.Context, dateString string,
	dateType entity.TaskDateType, weekNumber entity.WeekNumber,
) (Tasks, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	// 日付をtime.Timeに変換する
	date, err := convertToTimeTask(dateString, dateType, weekNumber)
	if err != nil {
		return nil, err
	}
	fmt.Println(user_id, *date, weekNumber)
	t := &entity.Task{
		UserID:     user_id,
		Date:       *date,
		WeekNumber: weekNumber,
	}
	ts, err := l.Repo.ListTasks(ctx, l.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	tasks := []Task{}
	for _, t := range ts {
		dateString := convertToStringTask(t.Date, t.DateType)
		tasks = append(tasks, Task{
			ID:         t.ID,
			Title:      t.Title,
			Date:       *dateString,
			DateType:   t.DateType,
			WeekNumber: t.WeekNumber,
		})
	}

	return tasks, nil
}

func convertToStringTask(date time.Time, dateType entity.TaskDateType) *string {
	// time.Timeの日付を文字列に変換する
	var dateString string
	switch dateType {
	case "Monthly":
		dateString = date.Format(time.RFC3339Nano)[0:7]
	case "Yearly", "Weekly":
		dateString = date.Format(time.RFC3339Nano)[0:4]
	}
	return &dateString
}
