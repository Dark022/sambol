package actions

import (
	"fmt"
	"net/http"

	"github.com/Dark022/sambol/models"
)

func (as *ActionSuite) Test_NewTemplateHandler() {
	res := as.HTML("/template/new").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "<section class=\"container mt-5\">")
}

func (as *ActionSuite) Test_SaveTemplate() {
	templateTest := &models.Template{Title: "Test", Content: "Testing post request", Active: true, Private: false}
	res := as.HTML("/template/save").Post(templateTest)
	err := as.DB.First(templateTest)
	as.NoError(err)
	as.NotZero(templateTest.ID)
	as.Equal("Test", templateTest.Title)
	as.Equal("Testing post request", templateTest.Content)
	as.Equal(true, templateTest.Active)
	as.Equal(false, templateTest.Private)
	as.Equal("/", res.Location())
	as.Equal(http.StatusSeeOther, res.Code)
}

func (as *ActionSuite) Test_ShowTemplate() {
	templateTest := &models.Template{Title: "Test", Content: "Testing show request", Active: true, Private: false}
	as.HTML("/template/save").Post(templateTest)
	err := as.DB.First(templateTest)
	as.NoError(err)
	res := as.HTML(fmt.Sprintf("/template/%v/show/", templateTest.ID)).Get()

	as.Contains(res.Body.String(), "Test")
	as.Contains(res.Body.String(), "Testing show request")
	as.Contains(res.Body.String(), "Active: true")
	as.Contains(res.Body.String(), "Private: false")
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_EditTemplate() {
	templateTest := &models.Template{Title: "Test", Content: "Testing edit request", Active: false, Private: true}
	as.HTML("/template/save").Post(templateTest)
	err := as.DB.First(templateTest)
	as.NoError(err)
	res := as.HTML(fmt.Sprintf("/template/%v/edit/", templateTest.ID)).Get()

	as.Contains(res.Body.String(), "Testing edit request")
	as.Contains(res.Body.String(), "false")
	as.Contains(res.Body.String(), "true")
	as.Contains(res.Body.String(), "Test")
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_DeleteTemplate() {
	templateTest := &models.Template{Title: "Test", Content: "Testing delete request", Active: false, Private: true}
	as.HTML("/template/save").Post(templateTest)
	err := as.DB.First(templateTest)
	as.NoError(err)
	res := as.HTML(fmt.Sprintf("/template/%v/delete/", templateTest.ID)).Delete()

	err = as.DB.First(templateTest)
	as.Error(err)
	as.Equal("/", res.Location())
	as.Equal(http.StatusSeeOther, res.Code)
}
