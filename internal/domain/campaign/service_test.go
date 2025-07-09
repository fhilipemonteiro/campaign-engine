package campaign

import (
	"campaign-engine/internal/contract"
	internalErrors "campaign-engine/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) GetAll() ([]Campaign, error) {
	args := r.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "This is a test campaign content.",
		Emails:  []string{"user@mail.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockRepository
	id, err := service.Create(newCampaign)
	assert.NotNil(id, "Expected campaign ID to be returned, but got nil.")
	assert.Nil(err, "Expected no error when creating a campaign, but got: %v", err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign := contract.NewCampaign{
		Name:    "",
		Content: "This is a test campaign content.",
		Emails:  []string{"user@mail.com"},
	}
	id, err := service.Create(newCampaign)
	assert.Empty(id, "Expected campaign ID to be empty when validation fails, but got: %s", id)
	assert.NotNil(err, "Expected an error when creating a campaign with invalid data, but got nil.")
	assert.EqualError(err, "campaign name cannot be empty", "Expected error message to be 'campaign name cannot be empty', but got '%s'.", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaignsEqual(newCampaign, campaign)
	})).Return(nil)
	service.Repository = repositoryMock
	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositoryError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalErrors.ErrInternal)
	service.Repository = repositoryMock
	_, err := service.Create(newCampaign)
	assert.NotNil(err, "Expected an error when saving the campaign, but got nil.")
	assert.True(errors.Is(err, internalErrors.ErrInternal), "Expected error to be of type ErrInternal, but got: %v", err)
}

func campaignsEqual(expected contract.NewCampaign, actual *Campaign) bool {
	if actual.Name != expected.Name || actual.Content != expected.Content {
		return false
	}
	if len(actual.Contacts) != len(expected.Emails) {
		return false
	}
	for i, email := range expected.Emails {
		if actual.Contacts[i].Email != email {
			return false
		}
	}
	return true
}
