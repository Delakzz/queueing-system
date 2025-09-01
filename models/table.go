package models

import "strconv"

type Table struct {
	ID int
}

func NewTable(id int) Table {
	return Table{ID: id}
}

func (t *Table) ShowInfo() string {
	return "Table " + strconv.Itoa(t.ID)
}
