// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package querytest

import (
	"database/sql"
)

type Bar struct {
	A     int32
	Alias sql.NullString
}

type Foo struct {
	A    int32
	Name sql.NullString
}
