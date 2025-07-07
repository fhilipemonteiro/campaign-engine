package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Test Campaign"
	content = "This is a test campaign content."
	emails  = []string{"user1@mail.com", "user2@mail.com"}
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, emails)
	assert.Equal(name, campaign.Name, "Expected campaign name to be '%s', but got '%s'.", name, campaign.Name)
	assert.Equal(content, campaign.Content, "Expected campaign content to be '%s', but got '%s'.", content, campaign.Content)
	assert.Len(campaign.Contacts, 2, "Expected campaign to have 2 contacts, but got %d.", len(campaign.Contacts))
	assert.Equal(emails[0], campaign.Contacts[0].Email,
		"Expected first contact email to be '%s', but got '%s'.", emails[0], campaign.Contacts[0].Email)
	assert.Equal(emails[1], campaign.Contacts[1].Email,
		"Expected second contact email to be '%s', but got '%s'.", emails[1], campaign.Contacts[1].Email)
}

func Test_NewCampaign_IDIsNotNull(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, content, emails)
	assert.NotNil(campaign.ID, "Expected campaign ID to be set, but it was empty.")
}

func Test_NewCampaign_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Second)
	campaign, _ := NewCampaign(name, content, emails)
	assert.NotNil(campaign.CreatedAt, "Expected campaign CreatedAt to be set, but it was empty.")
	assert.Greater(campaign.CreatedAt, now, "Expected campaign CreatedAt to be set in the past, but it was empty.")
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, emails)
	assert.NotNil(err, "Expected an error when creating a campaign with an empty name, but got nil.")
	assert.EqualError(err, "campaign name cannot be empty", "Expected error message to be 'campaign name cannot be empty', but got '%s'.", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", emails)
	assert.NotNil(err, "Expected an error when creating a campaign with empty content, but got nil.")
	assert.EqualError(err, "campaign content cannot be empty", "Expected error message to be 'campaign content cannot be empty', but got '%s'.", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{})
	assert.NotNil(err, "Expected an error when creating a campaign with no contacts, but got nil.")
	assert.EqualError(err, "campaign must have at least one contact", "Expected error message to be 'campaign must have at least one contact', but got '%s'.", err.Error())
}

func Test_NewCampaign_MustValidateContactsEmails(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{""})
	assert.NotNil(err, "Expected an error when creating a campaign with an empty contact email, but got nil.")
	assert.EqualError(err, "contact email cannot be empty", "Expected error message to be 'contact email cannot be empty', but got '%s'.", err.Error())
}

func Test_NewCampaign_MustValidateContactsEmails_IsValidEmail(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{"mail.com"})
	assert.NotNil(err, "Expected an error when creating a campaign with an invalid contact email, but got nil.")
	assert.EqualError(err, "invalid email format: mail.com", "Expected error message to be 'invalid email format: mail.com', but got '%s'.", err.Error())
}
