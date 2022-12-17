package store

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
)

func (r *Repository) EditWish(
	ctx context.Context, db Execer, w *entity.Wish,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	w.Modified = r.Clocker.Now()
	sql := `UPDATE wish SET
		content = ?, modified = ?
	WHERE id = ?`
	result, err := db.ExecContext(
		ctx, sql, w.Content, w.Modified, w.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	w.ID = entity.WishID(id)
	return nil
}

func (r *Repository) CreateWish(
	ctx context.Context, db Execer, w *entity.Wish,
) error {
	w.Created = r.Clocker.Now()
	w.Modified = r.Clocker.Now()
	sql := `INSERT INTO wish(
		user_id, content, created, modified
	)
	VALUES (?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, w.UserID, w.Content, w.Created, w.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	w.ID = entity.WishID(id)
	return nil
}

func (r *Repository) FetchWish(
	ctx context.Context, db Queryer, c *entity.Wish,
) (entity.Wishs, error) {
	wishes := entity.Wishs{}
	sql := `SELECT
				id, user_id, content,
				created, modified
			FROM wish
			WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &wishes, sql, c.UserID); err != nil {
		return nil, err
	}
	return wishes, nil
}
