package models

func (ms *ModelSuite) Test_TableLoad() {
	testTemplates := []Template{
		{Title: "Title1", Content: "Content1", Active: false, Private: false},
		{Title: "Title2", Content: "Content2", Active: true, Private: true},
		{Title: "Title3", Content: "Content3", Active: false, Private: false},
		{Title: "Title4", Content: "Content4", Active: true, Private: true},
		{Title: "Title5", Content: "Content5", Active: false, Private: false},
	}

	for _, template := range testTemplates {
		ms.DB.Create(&template)
	}

	allTemplates := LoadTable()

	for i, template := range allTemplates {
		ms.Equal(testTemplates[i].Title, template.Title)
		ms.Equal(testTemplates[i].Content, template.Content)
		ms.Equal(testTemplates[i].Active, template.Active)
		ms.Equal(testTemplates[i].Private, template.Private)
	}
}

/*func (ms *ModelSuite) Test_ViewValidation() {
	testTemplates := []struct {
		Title     string
		Content   string
		ExpectedT string
		ExpectedC string
	}{
		{" ", " ", "empty", "empty"},
		{"Test", " ", "fill", "empty"},
		{" ", "Test", "empty", "fill"},
		{"Title1", " ", "same_title", "empty"},
		{"Title1", "Test", "same_title", "fill"},
		{"Test", "Test", "fill", "fill"},
	}

	template := Template{Title: "Title1", Content: "Content1", Active: false, Private: false}
	ms.DB.Create(&template)

	for _, template := range testTemplates {
		testTemplate := Template{Title: template.Title, Content: template.Content, Active: false, Private: false}
		titleResult, contentResult := ViewValidation(testTemplate)
		ms.Equal(titleResult, template.ExpectedT)
		ms.Equal(contentResult, template.ExpectedC)
	}
}*/

func (ms *ModelSuite) Test_SearchID() {
	template := Template{Title: "Title1", Content: "Content1", Active: false, Private: false}
	ms.DB.Create(&template)

	testTemplate := SearchID(template.ID)

	ms.Equal(testTemplate.ID, template.ID)
	ms.Equal(testTemplate.Title, template.Title)
	ms.Equal(testTemplate.Content, template.Content)
	ms.Equal(testTemplate.Active, template.Active)
	ms.Equal(testTemplate.Private, template.Private)
}

func (ms *ModelSuite) Test_DeleteRow() {
	template := Template{Title: "Title1", Content: "Content1", Active: false, Private: false}
	ms.DB.Create(&template)

	DeleteRow(template.ID)

	ms.Error(ms.DB.First(&template))
}
