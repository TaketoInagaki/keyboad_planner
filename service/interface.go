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

type TaskDeleter interface {
	DeleteTask(ctx context.Context, db store.Execer, t *entity.Task) error
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
	EditReflection(ctx context.Context, db store.Execer, r *entity.Reflection) error
	CreateReflection(ctx context.Context, db store.Execer, preDb store.Queryer, r *entity.Reflection) error
}

type ReflectionFetcher interface {
	FetchReflection(ctx context.Context, db store.Queryer, ref *entity.Reflection) (entity.Reflections, error)
}

type CheckCreator interface {
	EditCheck(ctx context.Context, db store.Execer, r *entity.Check) error
	CreateCheck(ctx context.Context, db store.Execer, r *entity.Check) error
}

type CheckFetcher interface {
	FetchCheck(ctx context.Context, db store.Queryer, ref *entity.Check) (entity.Checks, error)
}

type ActionCreator interface {
	EditAction(ctx context.Context, db store.Execer, a *entity.Action) error
	CreateAction(ctx context.Context, db store.Execer, a *entity.Action) error
}

type ActionUpdater interface {
	UpdateAction(ctx context.Context, db store.Execer, a *entity.Action) error
}

type ActionFetcher interface {
	FetchAction(ctx context.Context, db store.Queryer, a *entity.Action) (entity.Actions, error)
}

type ContinuationCreator interface {
	EditContinuation(ctx context.Context, db store.Execer, c *entity.Continuation) error
	CreateContinuation(ctx context.Context, db store.Execer, c *entity.Continuation) error
}

type ContinuationFetcher interface {
	FetchContinuation(ctx context.Context, db store.Queryer, c *entity.Continuation) (entity.Continuations, error)
}

type WishCreator interface {
	EditWish(ctx context.Context, db store.Execer, w *entity.Wish) error
	CreateWish(ctx context.Context, db store.Execer, w *entity.Wish) error
}

type WishFetcher interface {
	FetchWish(ctx context.Context, db store.Queryer, w *entity.Wish) (entity.Wishes, error)
}
