package services

import (
	"fmt"
	"queueing-system/models"
)

type Queue struct {
	capacity int
	current  int
	filePath string
}

func NewQueue(capacity int, filePath string) *Queue {
	return &Queue{
		capacity: capacity,
		current:  1,
		filePath: filePath,
	}
}

func (q *Queue) IsFull() bool {
	return q.capacity == q.current
}

func (q *Queue) NextNumber(table models.Table) string {
	var result string
	if q.IsFull() {
		q.current = 1
	}
	result = fmt.Sprintf("Customer %d, please proceed to Table %d", q.current, table.ID)
	q.increment()
	return result
}

func (q *Queue) increment() {
	q.current++
}
