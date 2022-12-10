package entity

import "time"

type ReflectionID int64
type ContentType string
type DateType string
type WeekNumber int16

const (
	ContentTypeChech  ContentType = "Chech"
	ContentTypeAction ContentType = "Action"
	ContentTypeNote   ContentType = "Note"
)

const (
	DateTypeDaily   DateType = "Daily"
	DateTypeWeekly  DateType = "Weekly"
	DateTypeMonthly DateType = "Monthly"
	DateTypeYearly  DateType = "Yearly"
)

type Reflection struct {
	ID          ReflectionID `json:"id" db:"id"`
	UserID      UserID       `json:"user_id" db:"user_id"`
	Content     string       `json:"content" db:"content"`
	ContentType ContentType  `json:"content_type" db:"content_type"`
	Date        time.Time    `json:"date" db:"date"`
	DateType    DateType     `json:"date_type" db:"date_type"`
	WeekNumber  WeekNumber   `json:"week_number" db:"week_number"`
	Created     time.Time    `json:"created" db:"created"`
	Modified    time.Time    `json:"modified" db:"modified"`
}

type Reflections []*Reflection
