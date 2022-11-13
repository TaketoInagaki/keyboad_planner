package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/service"
	"github.com/TaketoInagaki/keyboard_planner/testutil"
)

func TestListTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		tasks []service.Task
		want  want
	}{
		"ok": {
			tasks: []service.Task{
				{
					ID:         1,
					Title:      "test1",
					Date:       "2022",
					DateType:   "Weekly",
					WeekNumber: 3,
				},
				{
					ID:         2,
					Title:      "test2",
					Date:       "2022-06",
					DateType:   "Monthly",
					WeekNumber: 0,
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/ok_rsp.json.golden",
			},
		},
		"empty": {
			tasks: []service.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			moq := &ListTasksServiceMock{}
			moq.ListTasksFunc = func(
				ctx context.Context, date string, dateType entity.TaskDateType,
				weekNumber entity.WeekNumber,
			) (service.Tasks, error) {
				if tt.tasks != nil {
					return tt.tasks, nil
				}
				return nil, errors.New("error from mock")
			}
			sut := ListTask{Service: moq}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
