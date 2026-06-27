package models

import "time"

type Receipt struct {
	Charges   float64
	EntryTime time.Time
	ExitTime  time.Time
	Duration  time.Duration
}
