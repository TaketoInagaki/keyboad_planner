package handler

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/service"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService RegisterUserService LoginService

// タスク
type ListTasksService interface {
	ListTasks(
		ctx context.Context, date string, dateType entity.TaskDateType,
		weekNumber entity.WeekNumber,
	) (service.Tasks, error)
}

type AddTaskService interface {
	CreateOrEditTask(
		ctx context.Context, id entity.TaskID, title string,
		date string, dateType entity.TaskDateType, weekNumber entity.WeekNumber,
	) (*entity.Task, error)
}

type DeleteTaskService interface {
	DeleteTask(
		ctx context.Context, id entity.TaskID,
	) (*entity.Task, error)
}

// ユーザー
type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}

// 振り返り
type CreateOrEditReflectionService interface {
	CreateOrEditReflection(
		ctx context.Context, id entity.ReflectionID, content string,
		date string, dateType entity.DateType, weekNumber entity.WeekNumber,
	) (*entity.Reflection, error)
}

type FetchReflectionService interface {
	FetchReflection(
		ctx context.Context, date string, dateType entity.DateType,
		weekNumber entity.WeekNumber,
	) (service.Reflections, error)
}

type CreateOrEditCheckService interface {
	CreateOrEditCheck(
		ctx context.Context, id entity.CheckID, content string,
		date string, dateType entity.DateType, weekNumber entity.WeekNumber,
	) (*entity.Check, error)
}
type FetchCheckService interface {
	FetchCheck(
		ctx context.Context, date string, dateType entity.DateType,
		weekNumber entity.WeekNumber,
	) (service.Checks, error)
}

type CreateOrEditActionService interface {
	CreateOrEditAction(
		ctx context.Context, id entity.ActionID, content string,
		date string, dateType entity.DateType, weekNumber entity.WeekNumber,
	) (*entity.Action, error)
}

type FetchActionService interface {
	FetchAction(
		ctx context.Context, date string, dateType entity.DateType,
		weekNumber entity.WeekNumber,
	) (service.Actions, error)
}

// 継続リスト
type CreateOrEditContinuationService interface {
	CreateOrEditContinuationList(
		ctx context.Context, id entity.ContinuationID, content string,
		contentType entity.ContinuationType,
	) (*entity.Continuation, error)
}

type FetchContinuationService interface {
	FetchContinuationList(ctx context.Context) (*service.Continuations, error)
}

type CreateOrEditWishService interface {
	CreateOrEditWishList(
		ctx context.Context, id entity.WishID, content string,
	) (*entity.Wish, error)
}

type FetchWishService interface {
	FetchWishList(ctx context.Context) (service.Wishes, error)
}
