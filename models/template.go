package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Template is the struct to represent an email template
type Template struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Title   string    `json:"title" db:"title"`
	Content string    `json:"content" db:"content"`
	Active  bool      `json:"active" db:"active"`
	Private bool      `json:"private" db:"private"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

func LoadTable() []Template {
	templates := []Template{}
	if err := DB.All(&templates); err != nil {
		log.Fatal(err)
	}
	return templates
}

func ViewValidation(template Template) (string, string) {
	templates := []Template{}
	if err := DB.All(&templates); err != nil {
		log.Fatal(err)
	}
	// 2 errores: content, title
	var titleValidation, contentValidation string
	Errors := validate.Validate(
		&validators.StringIsPresent{Field: template.Title, Name: template.Title, Message: "Title can't be blank"},
		&validators.StringIsPresent{Field: template.Content, Name: template.Content, Message: "Content can't be blank"},
	)
	errors := Errors.Errors
	if len(errors[""]) == 2 {
		titleValidation = "empty"
		contentValidation = "empty"
	} else if len(errors[""]) == 1 {
		if strings.Contains(errors[""][0], "Content") {
			titleValidation = "fill"
			contentValidation = "empty"
		} else if strings.Contains(errors[""][0], "Title") {
			titleValidation = "empty"
			contentValidation = "fill"
		}
	} else {
		titleValidation = "fill"
		contentValidation = "fill"
	}

	for _, tmpt := range templates {
		if strings.ToLower(tmpt.Title) == strings.ToLower(template.Title) {
			titleValidation = "same_title"
		}
	}

	return titleValidation, contentValidation
}

func SearchID(id uuid.UUID) Template {
	template := Template{}
	if err := DB.Find(&template, id); err != nil {
		log.Fatal(err)
	}

	return template
}

func DeleteRow(id uuid.UUID) {
	template := &Template{ID: id}
	if err := DB.Destroy(template); err != nil {
		fmt.Println(err)
	}
}
