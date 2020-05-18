package models

import (
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
			err = tx.Create(&categories)
		}
	}

	return err
}
