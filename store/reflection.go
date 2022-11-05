package store

import (
	"context"

	"github.com/budougumi0617/go_todo_app/entity"
)

func (r *Repository) CreateOrEditReflection(
	ctx context.Context, db Execer, ref *entity.Reflection,
) error {
	ref.Created = r.Clocker.Now()
	ref.Modified = r.Clocker.Now()
	sql := `INSERT INTO reflection(
		      user_id, content, content_type, date,
			    date_type, created, modified
			  )
	    VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := db.ExecContext(
		ctx, sql, ref.UserID, ref.Content, ref.ContentType,
		ref.Date, ref.DateType, ref.Created, ref.Modified,
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

func (r *Repository) FetchReflections(
	ctx context.Context, db Queryer, id entity.UserID,
) (entity.Reflections, error) {
	reflections := entity.Reflections{}
	sql := `SELECT 
				id, user_id, content,
				content_type, date, date_type,
				created, modified 
			FROM reflection
			WHERE user_id = ?;`
	if err := db.SelectContext(ctx, &reflections, sql, id); err != nil {
		return nil, err
	}
	return reflections, nil
}
