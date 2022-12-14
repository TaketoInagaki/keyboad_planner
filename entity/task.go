package entity

import "time"

type TaskID int64
type TaskDateType string

const (
	TaskDateTypeWeekly  TaskDateType = "Weekly"
	TaskDateTypeMonthly TaskDateType = "Monthly"
	TaskDateTypeYearly  TaskDateType = "Yearly"
)

type Task struct {
	ID         TaskID       `json:"id" db:"id"`
	UserID     UserID       `json:"user_id" db:"user_id"`
	Title      string       `json:"title" db:"title"`
	Date       time.Time    `json:"date" db:"date"`
	DateType   TaskDateType `json:"date_type" db:"date_type"`
	WeekNumber WeekNumber   `json:"week_number" db:"week_number"`
	DeleteFlg  DeleteFlg    `json:"delete_flg" db:"delete_flg"`
	Created    time.Time    `json:"created" db:"created"`
	Modified   time.Time    `json:"modified" db:"modified"`
}

type Tasks []*Task
