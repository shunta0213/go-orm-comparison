// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package queries

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Age       interface{}
	CreatedAt time.Time
}
