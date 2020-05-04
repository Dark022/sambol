package models

import (
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"firstname" db:"firstname"`
	LastName  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"email" db:"email"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

func LoadUserTable() ([]User, error) {
	users := []User{}
	err := DB.All(&users)

	return users, err
}

func (u *User) UserValidation(tx *pop.Connection, id uuid.UUID) *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName", Message: "First Name can't be blank"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName", Message: "Last Name can't be blank"},
		&validators.EmailIsPresent{Field: u.Email, Name: "Email", Message: "Invalid Email"},

		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "Email",
			Message: "Email \"%v\" already registered",
			Fn: func() bool {
				emailExists, err := tx.Where("email = ?", u.Email).Where("id != ?", id).Exists(&User{})
				if err != nil {
					return false
				}
				return !emailExists
			},
		},
	)
}

func SearchUserID(id uuid.UUID) (User, error) {
	user := User{}
	err := DB.Find(&user, id)

	return user, err
}

func DeleteUser(id uuid.UUID) error {
	user := &User{ID: id}
	err := DB.Destroy(user)

	return err
}
