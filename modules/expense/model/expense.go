package model

import (
	"github.com/lib/pq"
)

type Expense struct {
	ID     int            `json:"id""`
	Title  string         `json:"title"`
	Amount float64        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `gorm:"type:string[]" json:"tags"`
}
