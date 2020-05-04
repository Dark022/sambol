package actions

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Dark022/sambol/models"
)

func (as *ActionSuite) Test_NewCampaign() {
	res := as.HTML("/campaign/new").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "<section class=\"container mt-5\">")
}

func (as *ActionSuite) Test_SaveCampaign() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"Test"},
		"Content": []string{"Testing"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	as.HTML("/template/save").Post(form)
	as.DB.First(templateTest)

	campaingTest := &models.Campaign{}
	formC := url.Values{
		"Name":       []string{"NameTest"},
		"StartDate":  []string{"2020-05-10"},
		"EndDate":    []string{"2020-06-10"},
		"TemplateID": []string{templateTest.ID.String()},
	}
	res := as.HTML("/campaign/save").Post(formC)

	as.NoError(as.DB.First(campaingTest))
	as.NotZero(campaingTest.ID)
	as.Equal("NameTest", campaingTest.Name)
	as.Equal(templateTest.ID, campaingTest.TemplateID)
	as.Contains(campaingTest.StartDate.String(), "2020-05-10")
	as.Contains(campaingTest.EndDate.String(), "2020-06-10")
	as.Equal("/", res.Location())
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_ShowCampaign() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"Test"},
		"Content": []string{"Testing"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	as.HTML("/template/save").Post(form)
	as.DB.First(templateTest)

	campaingTest := &models.Campaign{}
	formC := url.Values{
		"Name":       []string{"NameTest"},
		"StartDate":  []string{"2020-05-10"},
		"EndDate":    []string{"2020-06-10"},
		"TemplateID": []string{templateTest.ID.String()},
	}
	as.HTML("/campaign/save").Post(formC)

	as.NoError(as.DB.First(templateTest))
	res := as.HTML(fmt.Sprintf("/template/%v/show/", campaingTest.ID)).Get()

	as.Contains(res.Body.String(), "NameTest")
	as.Contains(res.Body.String(), "May 10, 2020")
	as.Contains(res.Body.String(), "Jun 06, 2020")
	as.Contains(res.Body.String(), templateTest.ID.String())
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_EditCampaign() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"TestSelect"},
		"Content": []string{"Testing"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	as.HTML("/template/save").Post(form)
	as.DB.First(templateTest)

	campaingTest := &models.Campaign{}
	formC := url.Values{
		"Name":       []string{"NameTest"},
		"StartDate":  []string{"2020-05-10"},
		"EndDate":    []string{"2020-06-10"},
		"TemplateID": []string{templateTest.ID.String()},
	}
	as.HTML("/campaign/save").Post(formC)
	as.NoError(as.DB.First(campaingTest))

	res := as.HTML(fmt.Sprintf("/template/%v/edit/", campaingTest.ID)).Get()

	as.Contains(res.Body.String(), "NameTest")
	as.Contains(res.Body.String(), "2020-05-10")
	as.Contains(res.Body.String(), "2020-06-10")
	as.Contains(res.Body.String(), "TestSelect")
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_DeleteCampaign() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"Test"},
		"Content": []string{"Testing"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	as.HTML("/template/save").Post(form)
	as.DB.First(templateTest)

	campaingTest := &models.Campaign{}
	formC := url.Values{
		"Name":       []string{"NameTest"},
		"StartDate":  []string{"2020-05-10"},
		"EndDate":    []string{"2020-06-10"},
		"TemplateID": []string{templateTest.ID.String()},
	}
	as.HTML("/campaign/save").Post(formC)
	as.DB.First(campaingTest)
	res := as.HTML(fmt.Sprintf("/template/%v/delete/", campaingTest.ID)).Delete()

	as.Error(as.DB.First(campaingTest))
	as.Equal("/", res.Location())
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_UpdateCampaign() {
	templateTest := &models.Template{Title: "OldTest", Content: "OldTest", Active: true, Private: false}
	as.HTML("/template/save").Post(templateTest)
	as.NoError(as.DB.First(templateTest))
	form := url.Values{
		"Title":   []string{"Test Update"},
		"Content": []string{"Testing update request"},
		"Active":  []string{"false"},
		"Private": []string{"true"},
	}
	ID = templateTest.ID.String()
	as.HTML(fmt.Sprintf("/template/%v/edit/", templateTest.ID)).Get()
	as.HTML("/template/edit/save").Put(form)
}
