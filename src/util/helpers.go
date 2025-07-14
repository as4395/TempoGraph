package util

import (
	"time"
)

// ParseTime parses an RFC3339 timestamp string to time.Time
func ParseTime(ts string) (time.Time, error) {
	return time.Parse(time.RFC3339, ts)
}
