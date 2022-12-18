package handler

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/service"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService RegisterUserService LoginService
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

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}

type LoginService interface {
	Login(ctx context.Context, name, pw string) (string, error)
}

type CreateOrEditReflectionService interface {
	CreateOrEditReflection(
		ctx context.Context, id entity.ReflectionID, content string,
		contentType entity.ReflectionType, date string,
		dateType entity.DateType, weekNumber entity.WeekNumber,
	) (*entity.Reflection, error)
}

type FetchReflectionService interface {
	FetchReflection(
		ctx context.Context, date string, dateType entity.DateType,
		weekNumber entity.WeekNumber,
	) (service.Reflections, error)
}

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
