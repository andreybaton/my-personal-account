package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

type BaseRepository[T any] struct {
	db        *sql.DB
	tableName string
}

func NewBaseRepository[T any](db *sql.DB, tableName string) *BaseRepository[T] {
	return &BaseRepository[T]{
		db:        db,
		tableName: tableName,
	}
}

func (r *BaseRepository[T]) GetDB() *sql.DB {
	return r.db
}

func (r *BaseRepository[T]) FindByID(ctx context.Context, id int) (*T, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", r.tableName)

	var entity T
	err := r.db.QueryRowContext(ctx, query, id).Scan(r.getScanFields(&entity)...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	return fmt.Errorf("method Create must be implemented in specific repository")
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	return fmt.Errorf("method Update must be implemented in specific repository")
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", r.tableName)
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *BaseRepository[T]) FindAll(ctx context.Context, limit, offset int) ([]*T, error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", r.tableName)

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*T
	for rows.Next() {
		var entity T
		if err := rows.Scan(r.getScanFields(&entity)...); err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}

	return entities, nil
}

func (r *BaseRepository[T]) Count(ctx context.Context) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", r.tableName)

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func (r *BaseRepository[T]) Exists(ctx context.Context, id int) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", r.tableName)

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	return exists, err
}

func (r *BaseRepository[T]) FindByField(ctx context.Context, field string, value interface{}) ([]*T, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", r.tableName, field)

	rows, err := r.db.QueryContext(ctx, query, value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*T
	for rows.Next() {
		var entity T
		if err := rows.Scan(r.getScanFields(&entity)...); err != nil {
			return nil, err
		}
		entities = append(entities, &entity)
	}

	return entities, nil
}

func (r *BaseRepository[T]) getScanFields(entity *T) []interface{} {
	val := reflect.ValueOf(entity).Elem()
	fields := make([]interface{}, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Field(i).Addr().Interface()
	}

	return fields
}
