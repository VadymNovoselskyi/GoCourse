package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/VadymNovoselskyi/GoCourse/notes"
)

func main() {
	title, err := getUserInput("Title: ")
	if err != nil {
		fmt.Print("Invalid input for title: ", err)
		return
	}

	content, err := getUserInput("Content: ")
	if err != nil {
		fmt.Print("Invalid input for content :", err)
		return
	}

	note, err := notes.New(title, content)
	if err != nil {
		fmt.Print("Error creating a note: ", err)
		return
	}

	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Printf("Error saving note: %v\n", err)
		return
	}
	fmt.Println("Saving the note succeeded")
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("Input is empty")
	}

	return input, nil

}

// import "github.com/VadymNovoselskyi/GoCourse/profit_calc"

// func main() {
// 	profit_calc.CalculateProfits()
// }
