package todo

import (
	"time"
)

// individual todo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// list for todo items
type List []item

func (l *List) Add() {
	// add item to todo list
}
