package main

import (
	"fmt"
	"queueing-system/models"
	"queueing-system/services"
)

func main() {
	q := services.NewQueue(60, "db/queue.db")
	t1 := models.Table{ID: 1}

	// simulate 25 people entering queue
	for i := 0; i < 25; i++ {
		num := q.NextNumber(t1)
		fmt.Print(num + "\n")
	}
}
