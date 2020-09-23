package domains

import (
	"time"
)

type TodoID string

type Todo struct {
	ID        TodoID
	Name      string
	CreatedAt time.Time
}
