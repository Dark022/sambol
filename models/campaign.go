package models

import (
	"strings"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

//Campaign is the struct to represent an campaing
type Campaign struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	TemplateID uuid.UUID `json:"template_id" db:"template_id"`
	StartDate  time.Time `json:"start_date" db:"start_date"`
	EndDate    time.Time `json:"end_date" db:"end_date"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

func LoadCampaignTable() ([]Campaign, error) {
	campaign := []Campaign{}
	err := DB.All(&campaign)

	return campaign, err
}

func (c *Campaign) CampaignValidation(tx *pop.Connection) *validate.Errors {

	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name", Message: "Name can't be blank"},
		&validators.TimeAfterTime{FirstName: "EndDate", FirstTime: c.EndDate, SecondName: "StartDate", SecondTime: c.StartDate, Message: "End Date needs to be after Start Date"},

		&validators.FuncValidator{
			Field:   c.Name,
			Name:    "Name",
			Message: "Name \"%v\" already registered",
			Fn: func() bool {
				nameExists, err := tx.Where("name = ?", c.Name).Exists(&Campaign{})
				if err != nil {
					return false
				}
				return !nameExists
			},
		},
	)
}

func SearchCampaignID(id uuid.UUID) (Campaign, error) {
	campaign := Campaign{}
	err := DB.Find(&campaign, id)

	return campaign, err
}

func TodayTomorrow() (string, string) {
	date := time.Now()
	date2 := date.AddDate(0, 0, 1)
	todayArr := strings.Split(date.String(), " ")
	tomorrowArr := strings.Split(date2.String(), " ")
	today := todayArr[0]
	tomorrow := tomorrowArr[0]

	return today, tomorrow
}

func DeleteCampaign(id uuid.UUID) error {
	campaign := &Campaign{ID: id}
	err := DB.Destroy(campaign)
	return err
}
