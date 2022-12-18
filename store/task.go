package store

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
)

func (r *Repository) EditTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `UPDATE task SET
		title = ?, date = ?, modified = ?
	WHERE id = ?`
	result, err := db.ExecContext(
		ctx, sql, t.Title, t.Date, t.Modified, t.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) UpdateTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	t.Modified = r.Clocker.Now()
	sql := `UPDATE task SET
		status = ?
	WHERE user_id = ?
		AND id = ?`
	result, err := db.ExecContext(
		ctx, sql, t.Status, t.UserID, t.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) CreateTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
			(user_id, title, date, date_type, week_number, created, modified)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, t.UserID, t.Title, t.Date,
		t.DateType, t.WeekNumber, t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer, t *entity.Task,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
				id, user_id, title, status, date, date_type,
				week_number, created, modified
			FROM task
			WHERE user_id = ?
				AND date = ?
				AND week_number = ?
				AND delete_flg = 0;`
	if err := db.SelectContext(ctx, &tasks, sql, t.UserID, t.Date, t.WeekNumber); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) DeleteTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	sql := `UPDATE task
			SET delete_flg = 1
			WHERE user_id = ?
				AND id = ?;`
	result, err := db.ExecContext(
		ctx, sql, t.UserID, t.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}
