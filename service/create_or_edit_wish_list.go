package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

type CreateOrEditWishList struct {
	DB   store.Execer
	Repo WishCreator
}

func (a *CreateOrEditWishList) CreateOrEditWishList(
	ctx context.Context, id entity.WishID, content string,
) (*entity.Wish, error) {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	con := &entity.Wish{
		ID:      id,
		UserID:  userId,
		Content: content,
	}
	if con.ID != 0 {
		if err := a.Repo.EditWish(ctx, a.DB, con); err != nil {
			return nil, fmt.Errorf("failed to edit: %w", err)
		}
	} else {
		if err := a.Repo.CreateWish(ctx, a.DB, con); err != nil {
			return nil, fmt.Errorf("failed to register: %w", err)
		}
	}
	return con, nil
}
