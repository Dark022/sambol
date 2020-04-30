package actions

import (
	"net/http"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

var ID string

func NewTemplate(c buffalo.Context) error {
	template := models.Template{}
	c.Set("template", template)
	return c.Render(http.StatusOK, r.HTML("templates/new.plush.html"))
}

func SaveTemplate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	template := models.Template{}

	if err := c.Bind(&template); err != nil {
		return err
	}

	//Validate if inputs are empty and if title is already registered
	if errors := template.ViewValidation(tx); errors.HasAny() {
		c.Set("template", template)
		c.Set("errors", errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("templates/new.plush.html"))
	}

	//Create table row
	if err := tx.Create(&template); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func ShowTemplate(c buffalo.Context) error {
	idStr := c.Param("template_id")

	id, err := uuid.FromString(idStr)
	if err != nil {
		return err
	}
	template := models.SearchID(id)
	c.Set("template", template)
	return c.Render(http.StatusOK, r.HTML("templates/show.plush.html"))
}

func DeleteTemplate(c buffalo.Context) error {
	idStr := c.Param("template_id")

	id, err := uuid.FromString(idStr)
	if err != nil {
		return err
	}
	models.DeleteRow(id)

	return c.Redirect(http.StatusSeeOther, "/")
}

func EditTemplate(c buffalo.Context) error {
	ID = c.Param("template_id")

	id, err := uuid.FromString(ID)
	if err != nil {
		return err
	}
	template := models.SearchID(id)
	c.Set("template", template)
	return c.Render(http.StatusOK, r.HTML("templates/edit.plush.html"))
}

func UpdateTemplate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	id, err := uuid.FromString(ID)
	if err != nil {
		return err
	}
	template := models.SearchID(id)

	templateForm := models.Template{}

	if err := c.Bind(&templateForm); err != nil {
		return err
	}

	//Validate if inputs are empty and if title is already registered
	if errors := template.ViewValidation(tx); errors.HasAny() {
		c.Set("template", template)
		c.Set("errors", errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("templates/new.plush.html"))
	}

	templateForm.ID = template.ID

	if err := tx.Update(&templateForm); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
