// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const schemaScopedCreate = `-- name: SchemaScopedCreate :execresult
INSERT INTO foo.bar (id, name) VALUES (?, ?)
`

type SchemaScopedCreateParams struct {
	ID   uint64
	Name string
}

func (q *Queries) SchemaScopedCreate(ctx context.Context, arg SchemaScopedCreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, schemaScopedCreate, arg.ID, arg.Name)
}
