package campaign

import (
	"errors"
	"regexp"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedAt time.Time
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("campaign name cannot be empty")
	}
	if content == "" {
		return nil, errors.New("campaign content cannot be empty")
	}
	if len(emails) == 0 {
		return nil, errors.New("campaign must have at least one contact")
	}
	if slices.Contains(emails, "") {
		return nil, errors.New("contact email cannot be empty")
	}
	for _, email := range emails {
		if !emailRegex.MatchString(email) {
			return nil, errors.New("invalid email format: " + email)
		}
	}
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}
	return &Campaign{
		ID:        uuid.NewString(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}, nil
}
