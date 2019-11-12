package model

import (
	"time"
)

// Assoc represets assoc
type Assoc struct {
	Date     string
	Distance float64
	PoS      int
	Entity   string
	Word     string
}

// DateFormat represents default date format
const DateFormat = time.RFC3339
