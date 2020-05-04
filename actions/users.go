package actions

import (
	"net/http"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func ListUser(c buffalo.Context) error {
	users, err := models.LoadUserTable()
	if err != nil {
		return err
	}
	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("users/list.plush.html"))
}

func NewUser(c buffalo.Context) error {
	user := models.Template{}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("users/new.plush.html"))
}

func SaveUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	//Validate if inputs are empty and if title is already registered
	if errors := user.UserValidation(tx, uuid.Nil); errors.HasAny() {
		c.Set("user", user)
		c.Set("errors", errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("users/new.plush.html"))
	}

	//Create table row
	if err := tx.Create(&user); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/user")
}

func ShowUser(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		return err
	}
	user, err := models.SearchUserID(id)
	if err != nil {
		return err
	}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("users/show.plush.html"))
}

func DeleteUser(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		return err
	}
	models.DeleteUser(id)

	return c.Redirect(http.StatusSeeOther, "/user")
}

func EditUser(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		return err
	}
	user, err := models.SearchUserID(id)
	if err != nil {
		return err
	}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("users/edit.plush.html"))
}

func UpdateUser(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	id, err := uuid.FromString(c.Param("user_id"))
	if err != nil {
		return err
	}
	user, err := models.SearchUserID(id)
	if err != nil {
		return err
	}

	userForm := models.User{}

	if err := c.Bind(&userForm); err != nil {
		return err
	}

	//Validate if inputs are empty and if title is already registered
	if errors := userForm.UserValidation(tx, id); errors.HasAny() {
		c.Set("user", user)
		c.Set("errors", errors)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("users/new.plush.html"))
	}

	userForm.ID = user.ID

	if err := tx.Update(&userForm); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/user")
}
