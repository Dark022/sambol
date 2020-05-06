package actions

import (
	"fmt"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/validate"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"options": func() map[string]string {

				templates, _ := models.LoadTable()
				options := make(map[string]string, len(templates))

				for _, template := range templates {
					options[template.Title] = template.ID.String()
				}
				return options
			},

			"errorIdentifier": func(errors *validate.Errors, err string) bool {
				for keys := range errors.Errors {
					if keys == err {
						return true
					}
				}
				return false
			},

			"userFullName": func(user models.User) string {
				return fmt.Sprintf("%v %v", user.FirstName, user.LastName)
			},
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
		},
	})
}
