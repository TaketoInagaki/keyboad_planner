package store

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
)

func (r *Repository) EditCheck(
	ctx context.Context, db Execer, c *entity.Check,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	c.Modified = r.Clocker.Now()
	sql := `UPDATE checklist SET
		content = ?, modified = ?
	WHERE id = ?`
	result, err := db.ExecContext(
		ctx, sql, c.Content, c.Modified, c.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = entity.CheckID(id)
	return nil
}

func (r *Repository) CreateCheck(
	ctx context.Context, db Execer, c *entity.Check,
) error {
	c.Created = r.Clocker.Now()
	c.Modified = r.Clocker.Now()
	sql := `INSERT INTO checklist(
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
	c.ID = entity.CheckID(id)
	return nil
}

func (r *Repository) FetchCheck(
	ctx context.Context, db Queryer, c *entity.Check,
) (entity.Checks, error) {
	checks := entity.Checks{}
	sql := `SELECT
				id, user_id, content,
				date, date_type,
				week_number, created, modified
			FROM checklist
			WHERE user_id = ?
				AND date = ?
				AND week_number = ?;`
	if err := db.SelectContext(ctx, &checks, sql, c.UserID, c.Date, c.WeekNumber); err != nil {
		return nil, err
	}
	return checks, nil
}

func (r *Repository) DeleteCheck(
	ctx context.Context, db Execer, c *entity.Check,
) error {
	sql := `UPDATE checklist
			SET delete_flg = 1
			WHERE user_id = ?
				AND id = ?;`
	result, err := db.ExecContext(
		ctx, sql, c.UserID, c.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = entity.CheckID(id)
	return nil
}
