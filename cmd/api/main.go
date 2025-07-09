package main

import (
	"campaign-engine/internal/domain/campaign"
	"campaign-engine/internal/endpoints"
	"campaign-engine/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{Repository: &database.CampaignRepository{}}
	handler := endpoints.Handler{CampaignService: campaignService}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignsGet)

	http.ListenAndServe(":8000", r)
}
