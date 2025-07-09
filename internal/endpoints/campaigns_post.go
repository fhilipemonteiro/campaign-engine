package endpoints

import (
	"campaign-engine/internal/contract"
	internalError "campaign-engine/internal/internal-errors"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var body contract.NewCampaign
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, http.StatusBadRequest, "INVALID_PAYLOAD", map[string]string{"message": err.Error()})
		return
	}
	id, err := h.CampaignService.Create(body)
	if err != nil {
		if errors.Is(err, internalError.ErrInternal) {
			writeJSON(w, http.StatusInternalServerError, "INTERNAL_ERROR", map[string]string{"message": err.Error()})
			return
		}
		writeJSON(w, http.StatusBadRequest, "INVALID_DATA", map[string]string{"message": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, "CAMPAIGN_CREATED", map[string]string{"id": id})
}
