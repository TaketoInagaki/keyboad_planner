package store

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
)

func (r *Repository) EditContinuation(
	ctx context.Context, db Execer, con *entity.Continuation,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	con.Modified = r.Clocker.Now()
	sql := `UPDATE continuation SET
		content = ?, modified = ?
	WHERE id = ?`
	result, err := db.ExecContext(
		ctx, sql, con.Content, con.Modified, con.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	con.ID = entity.ContinuationID(id)
	return nil
}

func (r *Repository) CreateContinuation(
	ctx context.Context, db Execer, con *entity.Continuation,
) error {
	con.Created = r.Clocker.Now()
	con.Modified = r.Clocker.Now()
	sql := `INSERT INTO continuation(
		user_id, content, content_type, created, modified
	)
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, con.UserID, con.Content, con.ContinuationType, con.Created, con.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	con.ID = entity.ContinuationID(id)
	return nil
}

func (r *Repository) FetchContinuation(
	ctx context.Context, db Queryer, c *entity.Continuation,
) (entity.Continuations, error) {
	continuations := entity.Continuations{}
	sql := `SELECT
				id, user_id, content,
				content_type, created, modified
			FROM continuation
			WHERE user_id = ?
				AND delete_flg = 0;`
	if err := db.SelectContext(ctx, &continuations, sql, c.UserID); err != nil {
		return nil, err
	}
	return continuations, nil
}

func (r *Repository) DeleteContinuation(
	ctx context.Context, db Execer, t *entity.Continuation,
) error {
	sql := `UPDATE continuation
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
	t.ID = entity.ContinuationID(id)
	return nil
}
