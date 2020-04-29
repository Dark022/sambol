package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/gobuffalo/uuid"
)

func Test_TableLoad(t *testing.T) {
	testTemplates := []struct {
		ID      uuid.UUID `json:"id" db:"id"`
		Title   string    `json:"title" db:"title"`
		Content string    `json:"content" db:"content"`
		Active  bool      `json:"active" db:"active"`
		Private bool      `json:"private" db:"private"`

		CreatedAt  time.Time `json:"created_at" db:"created_at"`
		UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
	}{
		{Title: "Title1", Content: "Content1", Active: false, Private: false},
		{Title: "Title2", Content: "Content2", Active: true, Private: true},
		{Title: "Title3", Content: "Content3", Active: false, Private: false},
		{Title: "Title4", Content: "Content4", Active: true, Private: true},
		{Title: "Title5", Content: "Content5", Active: false, Private: false},
	}

	for _, template := range testTemplates {
		DB.Create(&template)
	}

	allTemplates := LoadTable()

	fmt.Println(allTemplates)
	for _, template := range allTemplates {
		fmt.Println(template, "a")
	}
}
