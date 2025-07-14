package graph

import "time"

type EntityState struct {
	Timestamp time.Time
	Data      Mutation
	Deleted   bool
}
