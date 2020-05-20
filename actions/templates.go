package actions

import (
	"net/http"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

var ID string

func ListTemplate(c buffalo.Context) error {
	private := c.Param("private")

	if private == "true" {
		templates, err := models.LoadPrivateTemplateTable()
		if err != nil {
			return err
		}
		c.Set("templates", templates)
		c.Set("current", "")
		c.Set("current2", "")
		c.Set("current3", "current-link")
		return c.Render(http.StatusOK, r.HTML("templates/list.plush.html"))
	}

	if private == "false" {
		templates, err := models.LoadPublicTemplateTable()
		if err != nil {
			return err
		}
		c.Set("templates", templates)
		c.Set("current", "")
		c.Set("current2", "current-link")
		c.Set("current3", "")
		return c.Render(http.StatusOK, r.HTML("templates/list.plush.html"))
	}

	templates, err := models.LoadTable()
	if err != nil {
		return err
	}
	c.Set("templates", templates)
	c.Set("current", "current-link")
	c.Set("current2", "")
	c.Set("current3", "")
	return c.Render(http.StatusOK, r.HTML("templates/list.plush.html"))
}

func NewTemplate(c buffalo.Context) error {
	template := models.Template{}
	c.Set("template", template)
	c.Set("categories", "")
	return c.Render(http.StatusOK, r.HTML("templates/new.plush.html"))
}

func SaveTemplate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	template := models.Template{}

	categories := struct {
		Categories string
	}{}

	if err := c.Bind(&template); err != nil {
		return err
	}

	if err := c.Bind(&categories); err != nil {
		return err
	}

	//Validate if inputs are empty and if title is already registered
	if errors := template.ViewValidation(tx, uuid.Nil); errors.HasAny() {
		c.Set("template", template)
		c.Set("categories", categories.Categories)
		c.Set("errors", errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("templates/new.plush.html"))
	}

	//Create table row
	if err := tx.Create(&template); err != nil {
		return err
	}

	if err := models.NewCategories(categories.Categories, tx, template.ID); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/template")
}

func ShowTemplate(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("template_id"))
	if err != nil {
		return err
	}
	template, err := models.SearchID(id)
	if err != nil {
		return err
	}
	c.Set("template", template)
	return c.Render(http.StatusOK, r.HTML("templates/show.plush.html"))
}

func DeleteTemplate(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("template_id"))
	if err != nil {
		return err
	}
	models.DeleteRow(id)

	return c.Redirect(http.StatusSeeOther, "/template")
}

func EditTemplate(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("template_id"))
	if err != nil {
		return err
	}
	template, err := models.SearchID(id)
	if err != nil {
		return err
	}
	c.Set("template", template)
	return c.Render(http.StatusOK, r.HTML("templates/edit.plush.html"))
}

func UpdateTemplate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	id, err := uuid.FromString(c.Param("template_id"))
	if err != nil {
		return err
	}
	template, err := models.SearchID(id)
	if err != nil {
		return err
	}

	categories := struct {
		Categories string
	}{}

	if err := c.Bind(&categories); err != nil {
		return err
	}

	templateForm := models.Template{}

	if err := c.Bind(&templateForm); err != nil {
		return err
	}

	//Validate if inputs are empty and if title is already registered
	if errors := templateForm.ViewValidation(tx, id); errors.HasAny() {
		c.Set("template", template)
		c.Set("errors", errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("templates/edit.plush.html"))
	}

	templateForm.ID = template.ID

	if err := tx.Update(&templateForm); err != nil {
		return err
	}

	if err := models.UpdateCategories(categories.Categories, tx, template.ID); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/template")
}
