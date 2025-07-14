package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"tempograph/graph"
)

// In-memory append-only log
var (
	logMu sync.Mutex
	logs  []logEntry
)

type logEntry struct {
	Timestamp time.Time
	Mutation  graph.Mutation
	Method    string
}

func AppendEvent(ts time.Time, mut graph.Mutation, method string) error {
	logMu.Lock()
	defer logMu.Unlock()
	logs = append(logs, logEntry{ts, mut, method})
	return nil
}

func Snapshot(at time.Time) []graph.Mutation {
	logMu.Lock()
	defer logMu.Unlock()

	var snapshot []graph.Mutation
	for _, entry := range logs {
		if entry.Timestamp.After(at) {
			break
		}
		snapshot = append(snapshot, entry.Mutation)
	}
	return snapshot
}

func DiffEvents(t1, t2 time.Time) []graph.Mutation {
	logMu.Lock()
	defer logMu.Unlock()

	var changes []graph.Mutation
	for _, entry := range logs {
		if entry.Timestamp.After(t1) && entry.Timestamp.Before(t2) {
			changes = append(changes, entry.Mutation)
		}
	}
	return changes
}
