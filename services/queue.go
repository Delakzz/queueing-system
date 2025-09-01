package services

import (
	"fmt"
	"queueing-system/models"
)

type Queue struct {
	Name     string
	tables   []models.Table
	capacity int
	current  int
	// filePath string
}

func NewQueue(name string, capacity int, filePath string) *Queue {
	return &Queue{
		Name:     name,
		capacity: capacity,
		current:  0,
		// filePath: filePath,
	}
}

func (q *Queue) GetName() string {
	return q.Name
}

func (q *Queue) IsEmpty() bool {
	return q.current == 0
}

func (q *Queue) IsFull() bool {
	return q.current > q.capacity
}

func (q *Queue) NextNumber(table models.Table) string {
	var result string
	if q.IsFull() || q.IsEmpty() {
		q.current = 1
	}
	result = fmt.Sprintf("Customer %d, please proceed to Table %d", q.current, table.ID)
	q.increment()
	return result
}

func (q *Queue) increment() {
	q.current++
}

func (q *Queue) AddTable() string {
	tableID := len(q.tables) + 1
	// fmt.Println(tableID)
	table := models.NewTable(tableID)
	// fmt.Println(table.ID)
	q.tables = append(q.tables, table)
	return table.ShowInfo()
}

func (q *Queue) GetTable(index int) models.Table {
	for i, t := range q.tables {
		if i == index {
			return t
		}
	}
	return models.Table{}
}

func (q *Queue) GetTables() []models.Table {
	return q.tables
}

func (q *Queue) RemoveTable(table models.Table) {
	for i, t := range q.tables {
		if t.ID == table.ID {
			q.tables = append(q.tables[:i], q.tables[i+1:]...)
			break
		}
	}
}
