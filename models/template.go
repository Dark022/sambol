package models

import "github.com/gobuffalo/uuid"

type Template struct{
	ID uuid.UUID `json="id" db="id"`
	Title string `json="title" db="title"`
	Content string `json="content" db="content"`
	Active string`json="active" db="active"`
	Private string`json="private" db="private"`
}

func Test(title string, content string, active string, private string){
	fmt.Println(title, content, active, private)
}