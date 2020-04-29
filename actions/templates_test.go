package actions

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Dark022/sambol/models"
)

func (as *ActionSuite) Test_NewTemplateHandler() {
	res := as.HTML("/template/new").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "<section class=\"container mt-5\">")
}

func (as *ActionSuite) Test_SaveTemplate() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"TestTitle"},
		"Content": []string{"TestContent"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	res := as.HTML("/template/save").Post(form)

	as.NoError(as.DB.First(templateTest))
	as.NotZero(templateTest.ID)
	as.Equal("TestTitle", templateTest.Title)
	as.Equal("TestContent", templateTest.Content)
	as.True(templateTest.Active)
	as.False(templateTest.Private)
	as.Equal("/", res.Location())
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_ShowTemplate() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"Test Show"},
		"Content": []string{"Testing show request"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	as.HTML("/template/save").Post(form)
	err := as.DB.First(templateTest)
	as.NoError(err)
	res := as.HTML(fmt.Sprintf("/template/%v/show/", templateTest.ID)).Get()

	as.Contains(res.Body.String(), "Test Show")
	as.Contains(res.Body.String(), "Testing show request")
	as.Contains(res.Body.String(), "Active: true")
	as.Contains(res.Body.String(), "Private: false")
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_EditTemplate() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"Test Edit Title"},
		"Content": []string{"Testing edit request"},
		"Active":  []string{"false"},
		"Private": []string{"true"},
	}
	as.HTML("/template/save").Post(form)
	as.NoError(as.DB.First(templateTest))
	res := as.HTML(fmt.Sprintf("/template/%v/edit/", templateTest.ID)).Get()

	as.Contains(res.Body.String(), "Testing edit request")
	as.Contains(res.Body.String(), "false")
	as.Contains(res.Body.String(), "true")
	as.Contains(res.Body.String(), "Test Edit Title")
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_DeleteTemplate() {
	templateTest := &models.Template{}
	form := url.Values{
		"Title":   []string{"Test Delete"},
		"Content": []string{"Testing delete request"},
		"Active":  []string{"true"},
		"Private": []string{"false"},
	}
	as.HTML("/template/save").Post(form)
	as.NoError(as.DB.First(templateTest))
	res := as.HTML(fmt.Sprintf("/template/%v/delete/", templateTest.ID)).Delete()

	as.Error(as.DB.First(templateTest))
	as.Equal("/", res.Location())
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_UpdateTemplate() {
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
