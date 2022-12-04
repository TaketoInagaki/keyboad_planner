package service

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegister UserGetter TokenGenerator
type TaskCreator interface {
	EditTask(ctx context.Context, db store.Execer, t *entity.Task) error
	CreateTask(ctx context.Context, db store.Execer, t *entity.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer, t *entity.Task) (entity.Tasks, error)
}
type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}

type UserGetter interface {
	GetUser(ctx context.Context, db store.Queryer, name string) (*entity.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u entity.User) ([]byte, error)
}

type ReflectionCreator interface {
	EditReflection(ctx context.Context, db store.Execer, t *entity.Reflection) error
	CreateReflection(ctx context.Context, db store.Execer, t *entity.Reflection) error
}
