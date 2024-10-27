package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TopicsToUse []string `json:"topicsToUse"`
}

const TasksPerStudent int = 5

func main() {
	// Check if the file path and number are provided as a command-line argument
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <file_path> <number_in_a_group_list>")
	}
	filePath := os.Args[1]
	inputNumber, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Cannot convert ", os.Args[2], " as a number")
		return
	}

	// Reading the list of items from the file
	items := readListFromFile(filePath)

	// Use the input number to create a seed
	seed := int64(time.Now().Year() + inputNumber)

	// Shuffle the list based on the seed
	shuffledItems := shuffleTasksList(items, seed)

	// Slice only top N required tasks
	yourItems := shuffledItems[:TasksPerStudent]

	// Print the tasks list
	fmt.Println("Your tasks:")
	for i, item := range yourItems {
		fmt.Printf("Task #%d: %+v\n", i+1, item)
	}
}

func readListFromFile(filePath string) []Task {
	// Read the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// Parse the JSON file
	var items []Task
	err = json.Unmarshal(fileContent, &items)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	return items
}

// Shuffle the items based on a seed (the input number)
func shuffleTasksList(items []Task, seed int64) []Task {
	r := rand.New(rand.NewSource(seed)) // Seed the random number generator

	shuffledItems := make([]Task, len(items))
	copy(shuffledItems, items)

	// Shuffle the list using Fisher-Yates algorithm
	for i := len(shuffledItems) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		shuffledItems[i], shuffledItems[j] = shuffledItems[j], shuffledItems[i]
	}

	return shuffledItems
}
