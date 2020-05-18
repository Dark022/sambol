package models

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

func LoadTable() ([]Template, error) {
	templates := []Template{}
	err := DB.All(&templates)

	return templates, err
}

func LoadPublicTemplateTable() ([]Template, error) {
	templates := []Template{}
	err := DB.Where("private = ?", false).All(&templates)

	return templates, err
}

func LoadPrivateTemplateTable() ([]Template, error) {
	templates := []Template{}
	err := DB.Where("private = ?", true).All(&templates)

	return templates, err
}

func (tmp *Template) ViewValidation(tx *pop.Connection, id uuid.UUID) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: tmp.Title, Name: "Title", Message: "Title can't be blank"},
		&validators.StringIsPresent{Field: tmp.Content, Name: "Content", Message: "Content can't be blank"},
		&validators.StringIsPresent{Field: tmp.Subject, Name: "Subject", Message: "Subject can't be blank"},
		&validators.StringIsPresent{Field: tmp.SenderName, Name: "SenderName", Message: "Sender name can't be blank"},

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

func SearchCategories(id uuid.UUID) map[int]string {
	category := []Categories{}
	templateCategories := []TemplateCategories{}

	DB.Where("template_id = ?", id).All(&templateCategories)

	categories := make(map[int]string, len(templateCategories))

	for i := range templateCategories {
		DB.Where("id = ?", templateCategories[i].CategoryID).All(&category)
		categories[i] = category[i].Name
	}

	return categories
}
