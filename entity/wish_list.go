package entity

import "time"

type WishID int64
type WishType string

type Wish struct {
	ID        WishID    `json:"id" db:"id"`
	UserID    UserID    `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	DeleteFlg DeleteFlg `json:"delete_flg" db:"delete_flg"`
	Created   time.Time `json:"created" db:"created"`
	Modified  time.Time `json:"modified" db:"modified"`
}

type Wishs []*Wish
