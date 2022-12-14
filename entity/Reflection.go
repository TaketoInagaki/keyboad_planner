package entity

import "time"

type ReflectionID int64
type ReflectionType string
type DateType string
type WeekNumber int16

const (
	ReflectionTypeChech  ReflectionType = "Chech"
	ReflectionTypeAction ReflectionType = "Action"
	ReflectionTypeNote   ReflectionType = "Note"
)

const (
	DateTypeDaily   DateType = "Daily"
	DateTypeWeekly  DateType = "Weekly"
	DateTypeMonthly DateType = "Monthly"
	DateTypeYearly  DateType = "Yearly"
)

type Reflection struct {
	ID             ReflectionID   `json:"id" db:"id"`
	UserID         UserID         `json:"user_id" db:"user_id"`
	Content        string         `json:"content" db:"content"`
	ReflectionType ReflectionType `json:"content_type" db:"content_type"`
	Date           time.Time      `json:"date" db:"date"`
	DateType       DateType       `json:"date_type" db:"date_type"`
	WeekNumber     WeekNumber     `json:"week_number" db:"week_number"`
	DeleteFlg      DeleteFlg      `json:"delete_flg" db:"delete_flg"`
	Created        time.Time      `json:"created" db:"created"`
	Modified       time.Time      `json:"modified" db:"modified"`
}

type Reflections []*Reflection
