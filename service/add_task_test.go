package service

import (
	"context"
	"testing"
	"time"

	"github.com/TaketoInagaki/keyboard_planner/auth"

	"github.com/google/go-cmp/cmp"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/store"
)

func TestCreateTask_CreateTask(t *testing.T) {
	t.Parallel()

	wantUID := entity.UserID(10)
	wantTitle := "test title"
	wantDate, _ := time.Parse("2006-01", "2022-11")
	wantDateType := entity.TaskDateType("Monthly")
	wantWeekNumber := entity.WeekNumber(3)
	wantTask := &entity.Task{
		UserID:     wantUID,
		Title:      wantTitle,
		Date:       wantDate,
		DateType:   wantDateType,
		WeekNumber: wantWeekNumber,
	}
	type TaskAdderMockParameter struct {
		in  *entity.Task
		err error
	}
	tests := map[string]struct {
		title string
		want  *entity.Task
		taprm TaskAdderMockParameter
	}{
		"正常系": {
			title: wantTitle,
			want:  wantTask,
			taprm: TaskAdderMockParameter{
				in:  wantTask,
				err: nil,
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ctx := auth.SetUserID(context.Background(), 10)
			task := &entity.Task{
				ID:         1,
				Title:      tt.want.Title,
				Date:       tt.want.Date,
				DateType:   tt.want.DateType,
				WeekNumber: tt.want.WeekNumber,
			}

			moqDB := &ExecerMock{}
			moqRepo := &TaskCreatorMock{}
			moqRepo.CreateTaskFunc = func(pctx context.Context, db store.Execer, task *entity.Task) error {
				if ctx != pctx {
					t.Fatalf("not want context %v", pctx)
				}
				if db != moqDB {
					t.Fatalf("not want db %v", db)
				}
				if d := cmp.Diff(task, tt.taprm.in); len(d) != 0 {
					t.Fatalf("differs: (-got +want)\n%s", d)
				}
				return tt.taprm.err
			}
			a := &CreateTask{
				DB:   moqDB,
				Repo: moqRepo,
			}
			// Create
			cgot, cerr := a.CreateOrEditTask(ctx, 0, task.Title,
				"2022-11-11", task.DateType, task.WeekNumber)
			if cerr != nil {
				t.Fatalf("want no error, but got %v", cerr)
				return
			}
			if d := cmp.Diff(cgot, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
			// Edit
			egot, eerr := a.CreateOrEditTask(ctx, task.ID, task.Title,
				"2022-11-11", task.DateType, task.WeekNumber)
			if eerr != nil {
				t.Fatalf("want no error, but got %v", eerr)
				return
			}
			if d := cmp.Diff(egot, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
