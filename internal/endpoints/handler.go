package endpoints

import "campaign-engine/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
