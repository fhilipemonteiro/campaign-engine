package endpoints

import (
	internalError "campaign-engine/internal/internal-errors"
	"errors"
	"net/http"
)

func (h *Handler) CampaignsGet(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.CampaignService.GetCampaigns()
	if err != nil {
		if errors.Is(err, internalError.ErrInternal) {
			writeJSON(w, http.StatusInternalServerError, "INTERNAL_ERROR", map[string]string{"message": "failed to retrieve campaigns"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, "UNKNOWN_ERROR", map[string]string{"message": "an unknown error occurred"})
		return
	}
	if len(campaigns) == 0 {
		writeJSON(w, http.StatusNotFound, "NO_CAMPAIGNS_FOUND", map[string]string{"message": "no campaigns found"})
		return
	}
	writeJSON(w, http.StatusOK, "CAMPAIGNS_RETRIEVED", campaigns)
}
