package store

import (
	"context"

	"github.com/TaketoInagaki/keyboard_planner/entity"
)

func (r *Repository) EditReflection(
	ctx context.Context, db Execer, ref *entity.Reflection,
) error {
	// TODO: 指定したidのデータがない時にその旨を知らせる
	ref.Modified = r.Clocker.Now()
	sql := `UPDATE reflection SET
		content = ?, modified = ?
	WHERE id = ?`
	result, err := db.ExecContext(
		ctx, sql, ref.Content, ref.Modified, ref.ID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ref.ID = entity.ReflectionID(id)
	return nil
}

func (r *Repository) CreateReflection(
	ctx context.Context, db Execer, ref *entity.Reflection,
) error {
	ref.Created = r.Clocker.Now()
	ref.Modified = r.Clocker.Now()
	sql := `INSERT INTO reflection(
		user_id, content, content_type, date,
		date_type, week_number, created, modified
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, ref.UserID, ref.Content, ref.ReflectionType,
		ref.Date, ref.DateType, ref.WeekNumber, ref.Created, ref.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ref.ID = entity.ReflectionID(id)
	return nil
}

func (r *Repository) FetchReflection(
	ctx context.Context, db Queryer, ref *entity.Reflection,
) (entity.Reflections, error) {
	reflections := entity.Reflections{}
	sql := `SELECT
				id, user_id, content,
				content_type, date, date_type,
				week_number, created, modified
			FROM reflection
			WHERE user_id = ?
				AND date = ?
				AND week_number = ?;`
	if err := db.SelectContext(ctx, &reflections, sql, ref.UserID, ref.Date, ref.WeekNumber); err != nil {
		return nil, err
	}
	return reflections, nil
}
