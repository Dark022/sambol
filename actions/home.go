package actions

import (
	"net/http"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	templates, err := models.LoadTable()
	if err != nil {
		return err
	}
	c.Set("templates", templates)
	return c.Render(http.StatusOK, r.HTML("templates/list.plush.html"))
}
