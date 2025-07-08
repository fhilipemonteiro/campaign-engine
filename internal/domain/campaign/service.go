package campaign

import (
	"campaign-engine/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(data contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(data.Name, data.Content, data.Emails)
	if err != nil {
		return "", err
	}
	if err = s.Repository.Save(campaign); err != nil {
		return "", err
	}
	return campaign.ID, nil
}
