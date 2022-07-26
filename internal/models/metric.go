package models

import "time"

type Metric struct {
	Key   string
	Value int
	Time  time.Time
}
