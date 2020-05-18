package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Template is the struct to represent an email template
type Template struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	Active     bool      `json:"active" db:"active"`
	Private    bool      `json:"private" db:"private"`
	Subject    string    `json:"subject" db:"subject"`
	SenderName string    `json:"sender_name" db:"sender_name"`
	Owner      string    `json:"owner" db:"owner"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

//Categories is the struct to represent a categorie
type Categories struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

//CampaignUsers is the struct to represent association between Campaign and Users
type TemplateCategories struct {
	ID         uuid.UUID `json:"id" db:"id"`
	TemplateID uuid.UUID `json:"template_id" db:"template_id"`
	CategoryID uuid.UUID `json:"category_id" db:"category_id"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

//CampaignUsers is the struct to represent association between Campaign and Users
type CampaignUsers struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CampaignID uuid.UUID `json:"campaign_id" db:"campaign_id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

//Campaign is the struct to represent a campaing
type Campaign struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	TemplateID uuid.UUID `json:"template_id" db:"template_id"`
	StartDate  time.Time `json:"start_date" db:"start_date"`
	EndDate    time.Time `json:"end_date" db:"end_date"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}

//User is the struct to represent a user
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"firstname" db:"firstname"`
	LastName  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"email" db:"email"`

	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdataedAt time.Time `json:"updated_at" db:"updated_at"`
}
