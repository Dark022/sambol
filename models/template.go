package models

import (
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

func LoadTable() ([]Template, error) {
	templates := []Template{}
	err := DB.All(&templates)

	return templates, err
}

func (tmp *Template) ViewValidation(tx *pop.Connection, id uuid.UUID) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: tmp.Title, Name: "Title", Message: "Title can't be blank"},
		&validators.StringIsPresent{Field: tmp.Content, Name: "Content", Message: "Content can't be blank"},

		&validators.FuncValidator{
			Field:   tmp.Title,
			Name:    "Title",
			Message: "Title \"%v\" already registered",
			Fn: func() bool {
				titleExists, err := tx.Where("title = ?", tmp.Title).Where("id != ?", id).Exists(&Template{})
				if err != nil {
					return false
				}
				return !titleExists
			},
		},
	)
}

func SearchID(id uuid.UUID) (Template, error) {
	template := Template{}
	err := DB.Find(&template, id)

	return template, err
}

func DeleteRow(id uuid.UUID) error {
	template := &Template{ID: id}
	err := DB.Destroy(template)

	return err
}
