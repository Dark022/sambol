package models

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/uuid"
)

func NewCategories(ctg string, tx *pop.Connection, templateID uuid.UUID) error {
	var err error
	categories := Categories{}
	templateCategories := TemplateCategories{}
	if strings.TrimSpace(ctg) != "" {
		if strings.Contains(ctg, ",") {
			categoriesArr := strings.Split(ctg, ",")

			for i := range categoriesArr {
				categories.Name = categoriesArr[i]
				if err = tx.Create(&categories); err != nil {
					return err
				}
				templateCategories.TemplateID = templateID
				templateCategories.CategoryID = categories.ID
				if err = tx.Create(&templateCategories); err != nil {
					return err
				}

				categories = Categories{}
				templateCategories = TemplateCategories{}
			}
		} else {
			categories.Name = ctg
			if err = tx.Create(&categories); err != nil {
				return err
			}

			templateCategories.TemplateID = templateID
			templateCategories.CategoryID = categories.ID
			if err = tx.Create(&templateCategories); err != nil {
				return err
			}
		}
	}

	return err
}

func UpdateCategories(ctg string, tx *pop.Connection, templateID uuid.UUID) error {
	var err error

	exists, err := tx.Where("template_id = ?", templateID).Exists(&TemplateCategories{})
	if err != nil {
		return err
	}

	if exists {
		err := tx.RawQuery("DELETE FROM template_categories WHERE template_id = ?", templateID).Exec()
		if err != nil {
			return err
		}
	}

	fmt.Println(ctg)

	if strings.TrimSpace(ctg) != "" {
		NewCategories(ctg, tx, templateID)
	}

	return err
}
