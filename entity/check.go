package entity

import "time"

type CheckID int64
type CheckType string

type Check struct {
	ID         CheckID    `json:"id" db:"id"`
	UserID     UserID     `json:"user_id" db:"user_id"`
	Content    string     `json:"content" db:"content"`
	Date       time.Time  `json:"date" db:"date"`
	DateType   DateType   `json:"date_type" db:"date_type"`
	WeekNumber WeekNumber `json:"week_number" db:"week_number"`
	DeleteFlg  DeleteFlg  `json:"delete_flg" db:"delete_flg"`
	Created    time.Time  `json:"created" db:"created"`
	Modified   time.Time  `json:"modified" db:"modified"`
}

type Checks []*Check
