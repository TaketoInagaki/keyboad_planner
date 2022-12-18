package entity

import "time"

type ActionID int64
type ActionType string
type ActioinStatus int16

const (
	ActioinStatusNotStarted ActioinStatus = 0
	ActioinStatusDone       ActioinStatus = 1
)

type Action struct {
	ID         ActionID      `json:"id" db:"id"`
	UserID     UserID        `json:"user_id" db:"user_id"`
	Content    string        `json:"content" db:"content"`
	Status     ActioinStatus `json:"status" db:"status"`
	Date       time.Time     `json:"date" db:"date"`
	DateType   DateType      `json:"date_type" db:"date_type"`
	WeekNumber WeekNumber    `json:"week_number" db:"week_number"`
	DeleteFlg  DeleteFlg     `json:"delete_flg" db:"delete_flg"`
	Created    time.Time     `json:"created" db:"created"`
	Modified   time.Time     `json:"modified" db:"modified"`
}

type Actions []*Action
