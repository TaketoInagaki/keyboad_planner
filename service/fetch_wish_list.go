package service

import (
	"context"
	"fmt"
	"time"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type FetchWishList struct {
	DB   store.Queryer
	Repo WishFetcher
}

type Wish struct {
	ID       entity.WishID `json:"id"`
	Content  string        `json:"content"`
	Created  string        `json:"created"`
	Modified string        `json:"modified"`
}

type Wishes []Wish

func (f *FetchWishList) FetchWishList(
	ctx context.Context,
) (Wishes, error) {
	user_id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	w := &entity.Wish{
		UserID: user_id,
	}
	// storeを使ってデータを取得する
	ws, err := f.Repo.FetchWish(ctx, f.DB, w)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}

	wishes := []Wish{}
	for _, w := range ws {
		created, modified := convertToStringWish(w.Created, w.Modified)
		wishes = append(wishes, Wish{
			ID:       w.ID,
			Content:  w.Content,
			Created:  *created,
			Modified: *modified,
		})
	}

	return wishes, nil
}

func convertToStringWish(createdTime time.Time, modifiedTime time.Time) (*string, *string) {
	// time.Timeの日付を文字列に変換する
	created := createdTime.Format(time.RFC3339Nano)[0:16]
	modified := modifiedTime.Format(time.RFC3339Nano)[0:16]
	return &created, &modified
}
