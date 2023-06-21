package model

import (
	"time"
)

type Entry struct {
	Value string
	Date  time.Time
}

func (e Entry) String() string {

	return e.Value
}
