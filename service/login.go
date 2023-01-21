package service

import (
	"context"
	"fmt"

	"github.com/TaketoInagaki/keyboard_planner/store"
)

type Login struct {
	DB             store.Queryer
	Repo           UserGetter
	TokenGenerator TokenGenerator
}

func (l *Login) Login(ctx context.Context, name, pw string) (string, int, error) {
	u, err := l.Repo.GetUser(ctx, l.DB, name)
	if err != nil {
		return "", 0, fmt.Errorf("failed to list: %w", err)
	}
	if err := u.ComparePassword(pw); err != nil {
		return "", 0, fmt.Errorf("wrong password: %w", err)
	}
	jwt, err := l.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", 0, fmt.Errorf("failed to generate JWT: %w", err)
	}

	return string(jwt), 7776000, nil
}
