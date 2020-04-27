package actions

import (
	"net/http"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo"
)

func AddTemplateForm(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("templates/new.plush.html"))
}

func AddTemplate(c buffalo.Context) error {
	t := &models.Template{}
	if err := c.Bind(t); err != nil {
		return err
	}
	tle := t.Title
	cnt := t.Content
	act := t.Active
	pvt := t.Private

	models.Test(tle, cnt, act, pvt)

	return c.Render(http.StatusOK, r.HTML("templates/list.plush.html"))
}
