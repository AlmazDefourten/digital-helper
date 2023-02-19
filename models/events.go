package models

import (
	"time"
)

type Grant struct {
	ID         int       `json:"id"`
	Names      string    `json:"name"`
	DateStarts time.Time `json:"dateStart"`
	DateEnds   time.Time `json:"dateEnd"`
}
