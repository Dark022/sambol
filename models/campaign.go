package models

import (
	"strings"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

func LoadCampaignTable() ([]Campaign, error) {
	campaign := []Campaign{}
	err := DB.All(&campaign)

	return campaign, err
}

func LoadCampaignUsers(id uuid.UUID) ([]User, error) {
	var err error
	campaignUsers := []CampaignUsers{}
	users := []User{}

	err = DB.Where("campaign_id = ?", id).All(&campaignUsers)
	if err != nil {
		return users, err
	}

	for i := range campaignUsers {
		err = DB.Where("id = ?", campaignUsers[i].UserID).All(&users)
		if err != nil {
			return users, err
		}
	}

	return users, err
}

func (c *Campaign) CampaignValidation(tx *pop.Connection, id uuid.UUID, users int) *validate.Errors {

	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name", Message: "Name can't be blank"},
		&validators.TimeAfterTime{FirstName: "EndDate", FirstTime: c.EndDate, SecondName: "StartDate", SecondTime: c.StartDate, Message: "End Date needs to be after Start Date"},
		&validators.IntIsPresent{Field: users, Name: "NoUserSelected", Message: "Please select at least one user"},

		&validators.FuncValidator{
			Field:   c.Name,
			Name:    "Name",
			Message: "Name \"%v\" already registered",
			Fn: func() bool {
				nameExists, err := tx.Where("name = ?", c.Name).Where("id != ?", id).Exists(&Campaign{})
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
