package main

import (
	"fmt"
	"queueing-system/models"
	"queueing-system/services"
	"queueing-system/utils"
)

func main() {
	q := services.NewQueue("Bank Queue", 50, "data/queue.json")

	for {
		showMainMenu()
		choice := utils.ReadInt("Enter your choice: ")
		fmt.Println()

		switch choice {
		case 1:
			addTable(q)
		case 2:
			table := updateQueue(q)
			showLatestCustomer(q, table)
		case 3:
			fmt.Print("Exiting...\n\n")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func showMainMenu() {
	fmt.Println()
	fmt.Println("======= Main Menu =======")
	fmt.Println("\n [1] Add Table")
	fmt.Println(" [2] Update Queue")
	fmt.Println(" [3] Exit")
	fmt.Println("\n=========================")
	fmt.Println()
}

func showLatestCustomer(q *services.Queue, table models.Table) {
	if table.ID == 0 {
		return
	}
	fmt.Println(q.NextNumber(table))
}

func addTable(q *services.Queue) {
	table := q.AddTable()
	fmt.Println(table + " is successfully added!")
}

func displayTables(q *services.Queue) bool {
	if len(q.GetTables()) == 0 {
		fmt.Println("No tables available.")
		return false
	}
	fmt.Print("======== Tables ========\n\n")
	for i, t := range q.GetTables() {
		fmt.Printf("[%d] %s\n", i+1, t.ShowInfo())
	}
	fmt.Println("[0] Cancel")
	fmt.Println("\n========================")
	return true
}

func updateQueue(q *services.Queue) models.Table {
	if !displayTables(q) {
		return models.Table{}
	}
	var tableIndex int
	for {
		tableIndex = utils.ReadInt("\nSelect a table by number: ")
		fmt.Println()
		if tableIndex == 0 {
			fmt.Println("Operation cancelled.")
			return models.Table{}
		} else if tableIndex < 1 || tableIndex > len(q.GetTables()) {
			fmt.Println("Invalid table selection.")
			continue
		}
		break
	}
	return q.GetTable(tableIndex - 1)
}
