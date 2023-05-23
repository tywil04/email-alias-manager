package domains

import (
	"encoding/json"
	"net/http"

	"EmailAliasManager/sharedlib/cryptography"
	"EmailAliasManager/sharedlib/db"

	"github.com/google/uuid"
)

type PostRequest struct {
	Domain string `json:"domain"`
}

type PostResponse struct {
	ID string `json:"id"`
}

type PatchRequest struct {
	DomainID string `json:"domainId"`
	Disabled bool   `json:"disabled"`
}

type PatchResponse struct {
	ID string `json:"id"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	requestData := PostRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	if requestData.Domain == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	domain := db.Domain{
		ID:     cryptography.RandomUUID(),
		User:   authedUser,
		Domain: requestData.Domain,
	}
	db.DB.Create(&domain)

	response := PostResponse{
		ID: domain.ID.String(),
	}
	json.NewEncoder(w).Encode(&response)
}

func Patch(w http.ResponseWriter, r *http.Request) {
	requestData := PatchRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	if requestData.DomainID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	domainId, err := uuid.Parse(requestData.DomainID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	domain := db.Domain{}
	db.DB.Find(&domain, db.Domain{ID: domainId, User: authedUser})
	db.DB.Model(&domain).Update("Disabled", requestData.Disabled)

	response := PatchResponse{
		ID: domain.ID.String(),
	}
	json.NewEncoder(w).Encode(&response)
}
