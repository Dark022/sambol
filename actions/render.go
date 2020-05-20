package actions

import (
	"fmt"
	"strings"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/uuid"
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

			"optionUsers": func() map[string]string {

				users, _ := models.LoadUserTable()
				options := make(map[string]string, len(users))

				for _, user := range users {
					options[user.FirstName+" "+user.LastName] = user.ID.String()
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

			"getTotalTemplates": func() int {
				var i int
				templates, _ := models.LoadTable()

				for i = range templates {
					i++
				}

				return i
			},

			"getTemplateOwner": func(ownerID string) string {
				ID, _ := uuid.FromString(ownerID)
				user, _ := models.SearchUserID(ID)
				return fmt.Sprintf("%v %v", user.FirstName, user.LastName)
			},

			"getTemplateCategories": func(id uuid.UUID) string {
				var str strings.Builder
				categories := models.SearchCategories(id)
				for i := range categories {
					str.WriteString(categories[i] + " ")
				}

				return str.String()
			},

			"getFormTemplateCategories": func(id uuid.UUID) string {
				var str strings.Builder
				categories := models.SearchCategories(id)
				for i := range categories {
					str.WriteString(categories[i] + " ")
				}
				strCategories := str.String()
				strCategories = strings.TrimSpace(strCategories)
				arrCtg := strings.Split(strCategories, " ")
				strCategories = strings.Join(arrCtg, ",")
				return strCategories
			},

			"getCheckedUsers": func(IDs struct{ UsersID []uuid.UUID }, id uuid.UUID) bool {
				for i := range IDs.UsersID {
					if id == IDs.UsersID[i] {
						return true
					}
				}
				return false
			},
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
		},
	})
}
