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

	allTemplates, err := LoadTable()
	ms.NoError(err)

	for i, template := range allTemplates {
		ms.Equal(testTemplates[i].Title, template.Title)
		ms.Equal(testTemplates[i].Content, template.Content)
		ms.Equal(testTemplates[i].Active, template.Active)
		ms.Equal(testTemplates[i].Private, template.Private)
	}
}

func (ms *ModelSuite) Test_ViewValidationEmpty() {
	testTemplates := []struct {
		Title     string
		Content   string
		ExpectedT string
		ExpectedC string
	}{
		{" ", " ", "blank", "blank"},
		{"Test", " ", "", "blank"},
		{" ", "Test", "blank", ""},
		{"Test", "Test", "", ""},
	}

	for _, template := range testTemplates {
		testTemplate := Template{
			Title:   template.Title,
			Content: template.Content,
			Active:  false,
			Private: false,
		}
		errors := testTemplate.ViewValidation(DB)
		errorK := errors.Keys()
		errorE := errors.Errors
		if len(errorK) == 2 {
			ms.Contains(errorE["content"][0], template.ExpectedC)
			ms.Contains(errorE["title"][0], template.ExpectedT)
		}

		if len(errorK) == 1 {
			if errorK[0] == "content" {
				ms.Contains(errorE["content"][0], template.ExpectedC)
			}

			if errorK[0] == "title" {
				ms.Contains(errorE["title"][0], template.ExpectedT)
			}
		}
	}
}

func (ms *ModelSuite) Test_ViewValidationExists() {
	template := Template{
		Title:   "Title1",
		Content: "Content",
		Active:  true,
		Private: false,
	}
	ms.DB.Create(&template)

	templateTest := Template{
		Title:   "Title1",
		Content: "Testing if exists validation",
		Active:  false,
		Private: true,
	}

	errors := templateTest.ViewValidation(DB)

	ms.Error(errors)
	ms.Contains(errors.String(), "registered")
}

func (ms *ModelSuite) Test_SearchID() {
	template := Template{Title: "Title1", Content: "Content1", Active: false, Private: false}
	ms.DB.Create(&template)

	testTemplate, err := SearchID(template.ID)
	ms.NoError(err)

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
