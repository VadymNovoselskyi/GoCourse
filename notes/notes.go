package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note *Note) Display() {
	fmt.Println("----------------------------")
	fmt.Println("Title:", note.Title)
	fmt.Println("Content:", note.Content)
	fmt.Println("Created at:", note.CreatedAt)
	fmt.Println("----------------------------")
}

func (note *Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	data, err := json.Marshal(note)
	if err != nil {
		errorMessage := fmt.Sprintf("Error converthing note to json: %v", err)
		return errors.New(errorMessage)
	}

	return os.WriteFile(filename, data, 0644)
}

func New(title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Invalid input empty value for either title, string or both")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
