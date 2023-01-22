package store

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
)

func (r *Repository) EditAction(
	ctx context.Context, db Execer, c *entity.Action,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	c.Modified = r.Clocker.Now()
	sql := `UPDATE actionlist SET
		content = ?, modified = ?
	WHERE user_id = ?
		AND id = ?`
	result, err := db.ExecContext(
		ctx, sql, c.Content, c.Modified, c.UserID, c.ID,
	)
	if err != nil {
		return err
	}
	var _, e = result.LastInsertId()
	if e != nil {
		return err
	}
	c.ID = entity.ActionID(c.ID)
	return nil
}

func (r *Repository) UpdateAction(
	ctx context.Context, db Execer, c *entity.Action,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	c.Modified = r.Clocker.Now()
	sql := `UPDATE actionlist SET
		status = ?
	WHERE user_id = ?
		AND id = ?`
	result, err := db.ExecContext(
		ctx, sql, c.Status, c.UserID, c.ID,
	)
	if err != nil {
		return err
	}
	var _, e = result.LastInsertId()
	if e != nil {
		return err
	}
	c.ID = entity.ActionID(c.ID)
	return nil
}

func (r *Repository) CreateAction(
	ctx context.Context, db Execer, c *entity.Action,
) error {
	c.Created = r.Clocker.Now()
	c.Modified = r.Clocker.Now()
	sql := `INSERT INTO actionlist(
		user_id, content, date, date_type,
		week_number, created, modified
	)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, c.UserID, c.Content, c.Date,
		c.DateType, c.WeekNumber, c.Created, c.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = entity.ActionID(id)
	return nil
}

func (r *Repository) FetchAction(
	ctx context.Context, db Queryer, c *entity.Action,
) (entity.Actions, error) {
	checks := entity.Actions{}
	sql := `SELECT
				id, user_id, content,
				status, date, date_type,
				week_number, created, modified
			FROM actionlist
			WHERE user_id = ?
				AND date = ?
				AND week_number = ?
				AND delete_flg = 0;`
	if err := db.SelectContext(ctx, &checks, sql, c.UserID, c.Date, c.WeekNumber); err != nil {
		return nil, err
	}
	return checks, nil
}

func (r *Repository) DeleteAction(
	ctx context.Context, db Execer, t *entity.Action,
) error {
	sql := `UPDATE actionlist
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
	t.ID = entity.ActionID(id)
	return nil
}
