package entity

import "time"

type ReflectionID int64
type ContentType string
type DateType string

const (
	ContentTypeChech  ContentType = "Chech"
	ContentTypeAction ContentType = "Action"
	ContentTypeNote ContentType = "Note"
)

const (
	DateTypeDaily DateType = "Daily"
	DateTypeWeekly DateType = "Weekly"
	DateTypeMonthly DateType = "Monthly"
	DateTypeYearly DateType = "Yearly"
)

type Reflection struct {
	ID          ReflectionID `json:"id" db:"id"`
	UserID      UserID       `json:"user_id" db:"user_id"`
	Content     string       `json:"title" db:"title"`
	ContentType ContentType  `json:"status" db:"status"`
	Date        time.Time    `json:"date" db:"date"`
	DateType    DateType     `json:"date_type" db:"date_type"`
	Created     time.Time    `json:"created" db:"created"`
	Modified    time.Time    `json:"modified" db:"modified"`
}

type Reflections []*Reflection
