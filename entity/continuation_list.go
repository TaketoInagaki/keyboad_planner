package entity

import "time"

type ContinuationID int64
type ContinuationType string
type DeleteFlg int16

const (
	DeleteFlgNormal DeleteFlg = 0
	DeleteFlgDelete DeleteFlg = 1
)
const (
	ContinuationTypeContinue ContinuationType = "Continue"
	ContinuationTypeBegin    ContinuationType = "Begin"
	ContinuationTypeQuit     ContinuationType = "Quit"
)

type Continuation struct {
	ID               ContinuationID   `json:"id" db:"id"`
	UserID           UserID           `json:"user_id" db:"user_id"`
	Content          string           `json:"content" db:"content"`
	ContinuationType ContinuationType `json:"content_type" db:"content_type"`
	DeleteFlg        DeleteFlg        `json:"delete_flg" db:"delete_flg"`
	Created          time.Time        `json:"created" db:"created"`
	Modified         time.Time        `json:"modified" db:"modified"`
}

type Continuations []*Continuation
