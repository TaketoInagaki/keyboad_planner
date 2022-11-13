package store

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TaketoInagaki/keyboard_planner/clock"
	"github.com/TaketoInagaki/keyboard_planner/entity"
	"github.com/TaketoInagaki/keyboard_planner/testutil"
	"github.com/TaketoInagaki/keyboard_planner/testutil/fixture"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
)

func prepareUser(ctx context.Context, t *testing.T, db Execer) entity.UserID {
	t.Helper()
	u := fixture.User(nil)
	result, err := db.ExecContext(ctx, "INSERT INTO user (name, password, role, created, modified) VALUES (?, ?, ?, ?, ?);", u.Name, u.Password, u.Role, u.Created, u.Modified)
	if err != nil {
		t.Fatalf("insert user: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("got user_id: %v", err)
	}
	return entity.UserID(id)
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) (*entity.Tasks) {
	t.Helper()
	userID := prepareUser(ctx, t, con)
	otherUserID := prepareUser(ctx, t, con)
	monthlyDate, _ := time.Parse("2006-01", "2022-11")
	yearlyWeeklyDate, _ := time.Parse("2006", "2022")
	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			UserID: userID,
			Title:  "want task 1", Date: monthlyDate,
			DateType: entity.TaskDateType("Monthly"),
			WeekNumber: entity.WeekNumber(0),
			Created: c.Now(), Modified: c.Now(),
		},
		{
			UserID: userID,
			Title:  "want task 2", Date: yearlyWeeklyDate,
			DateType: entity.TaskDateType("Yearly"),
			WeekNumber: entity.WeekNumber(0),
			Created: c.Now(), Modified: c.Now(),
		},
	}
	tasks := entity.Tasks{
		wants[0],
		{
			UserID: otherUserID,
			Title:  "not want task", Date: yearlyWeeklyDate,
			DateType: entity.TaskDateType("Weekly"),
			WeekNumber: entity.WeekNumber(9),
			Created: c.Now(), Modified: c.Now(),
		},
		wants[1],
	}
	result, err := con.ExecContext(ctx,
		`INSERT INTO task (user_id, title, date, date_type, week_number, created, modified)
			VALUES
			    (?, ?, ?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?, ?, ?);`,
		tasks[0].UserID, tasks[0].Title, tasks[0].Date, tasks[0].DateType, tasks[0].WeekNumber, tasks[0].Created, tasks[0].Modified,
		tasks[1].UserID, tasks[1].Title, tasks[1].Date, tasks[1].DateType, tasks[1].WeekNumber, tasks[1].Created, tasks[1].Modified,
		tasks[2].UserID, tasks[2].Title, tasks[2].Date, tasks[2].DateType, tasks[2].WeekNumber, tasks[2].Created, tasks[2].Modified,
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	tasks[0].ID = entity.TaskID(id)
	tasks[1].ID = entity.TaskID(id + 1)
	tasks[2].ID = entity.TaskID(id + 2)
	return &wants
}

func TestRepository_ListTasks(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	// entity.Taskを作成する他のテストケースと混ざるとテストがフェイルする。
	// そのため、トランザクションをはることでこのテストケースの中だけのテーブル状態にする。
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	monthlyDate, _ := time.Parse("2006-01", "2022-11")
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := prepareTasks(ctx, t, tx)
	want  := &entity.Task{
		UserID:     33,
		Date:       monthlyDate,
		WeekNumber: entity.WeekNumber(0),
	}

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx, want)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}

func TestRepository_AddTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	monthlyDate, _ := time.Parse("2006-01", "2022-11")

	c := clock.FixedClocker{}
	var wantID int64 = 20
	okTask := &entity.Task{
		UserID:     33,
		Title:      "ok task",
		Date:       monthlyDate,
		DateType:   entity.TaskDateType("Monthly"),
		WeekNumber: entity.WeekNumber(0),
		Created:    c.Now(),
		Modified:   c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectExec(
		// エスケープが必要
		`INSERT INTO task \(user_id, title, status, created, modified\) VALUES \(\?, \?, \?, \?, \?\)`,
	).WithArgs(okTask.UserID, okTask.Title, okTask.Date, okTask.DateType, okTask.WeekNumber, c.Now(), c.Now()).
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.CreateTask(ctx, xdb, okTask); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}
