package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gobuffalo/pop/v5"
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

func (tmp *Template) ViewValidation(tx *pop.Connection) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: tmp.Title, Name: "Title", Message: "Title can't be blank"},
		&validators.StringIsPresent{Field: tmp.Title, Name: "Content", Message: "Content can't be blank"},

		&validators.FuncValidator{
			Field:   tmp.Title,
			Name:    "Title",
			Message: "Title \"%v\" already registered",
			Fn: func() bool {
				titleExists, err := tx.Where("title = ?", tmp.Title).Exists(&Template{})
				if err != nil {
					return false
				}
				return !titleExists
			},
		},
	)
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
