package actions

import (
	"net/http"

	"github.com/Dark022/sambol/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

func ListCampaign(c buffalo.Context) error {

	campaign, err := models.LoadCampaignTable()
	if err != nil {
		return err
	}

	c.Set("campaigns", campaign)
	return c.Render(http.StatusOK, r.HTML("campaigns/list.plush.html"))
}

func NewCampaign(c buffalo.Context) error {
	today, tomorrow := models.TodayTomorrow()
	campaign := models.Campaign{}
	c.Set("campaign", campaign)
	c.Set("today", today)
	c.Set("tomorrow", tomorrow)

	return c.Render(http.StatusOK, r.HTML("campaigns/new.plush.html"))
}

func SaveCampaign(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	campaign := models.Campaign{}

	if err := c.Bind(&campaign); err != nil {
		return err
	}

	//Validate if inputs are empty and if name is already registered
	if errors := campaign.CampaignValidation(tx, uuid.Nil); errors.HasAny() {
		c.Set("campaign", campaign)
		c.Set("errors", errors)
		c.Set("today", campaign.StartDate)
		c.Set("tomorrow", campaign.EndDate)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("campaigns/new.plush.html"))
	}

	//Create table row
	if err := tx.Create(&campaign); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/campaign")
}

func ShowCampaign(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("campaign_id"))
	if err != nil {
		return err
	}
	campaign, err := models.SearchCampaignID(id)
	if err != nil {
		return err
	}

	c.Set("campaign", campaign)
	return c.Render(http.StatusOK, r.HTML("campaigns/show.plush.html"))
}

func DeleteCampaing(c buffalo.Context) error {
	id, err := uuid.FromString(c.Param("campaign_id"))
	if err != nil {
		return err
	}
	models.DeleteCampaign(id)

	return c.Redirect(http.StatusSeeOther, "/campaign")
}

func EditCampaign(c buffalo.Context) error {
	today, _ := models.TodayTomorrow()

	id, err := uuid.FromString(c.Param("campaign_id"))
	if err != nil {
		return err
	}
	campaign, err := models.SearchCampaignID(id)
	if err != nil {
		return err
	}
	c.Set("campaign", campaign)
	c.Set("today", today)
	c.Set("dateValue", campaign.StartDate)

	return c.Render(http.StatusOK, r.HTML("campaigns/edit.plush.html"))
}

func UpdateCampaign(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	id, err := uuid.FromString(c.Param("campaign_id"))
	if err != nil {
		return err
	}
	campaign, err := models.SearchCampaignID(id)
	if err != nil {
		return err
	}

	campaignForm := models.Campaign{}

	if err := c.Bind(&campaignForm); err != nil {
		return err
	}
	today, _ := models.TodayTomorrow()
	//Validate if inputs are empty and if title is already registered
	if errors := campaignForm.CampaignValidation(tx, id); errors.HasAny() {
		c.Set("campaign", campaign)
		c.Set("errors", errors)
		c.Set("today", today)
		c.Set("dateValue", campaignForm.StartDate)

		return c.Render(http.StatusUnprocessableEntity, r.HTML("campaigns/edit.plush.html"))
	}

	campaignForm.ID = campaign.ID

	if err := tx.Update(&campaignForm); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/campaign")
}
