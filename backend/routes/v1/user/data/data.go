package data

import (
	"encoding/json"
	"net/http"

	"EmailAliasManager/lib/db"

	"github.com/google/uuid"
)

type GetRequest struct{}

type GetResponse db.User

type PostRequest struct {
	Aliases []db.Alias
	Domains []db.Domain
}

type PostResponse struct{}

func Get(w http.ResponseWriter, r *http.Request) {
	requestData := GetRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	json.NewEncoder(w).Encode(&authedUser)
}

func Post(w http.ResponseWriter, r *http.Request) {
	requestData := PostRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	json.NewEncoder(w).Encode(&authedUser)

	for _, domain := range requestData.Domains {
		newId, _ := uuid.NewRandom()
		oldId := domain.ID
		domain.ID = newId
		domain.User = authedUser
		db.DB.Create(&domain)

		for _, alias := range requestData.Aliases {
			if alias.DomainID == oldId {
				newId, _ := uuid.NewRandom()
				alias.ID = newId
				alias.User = authedUser
				alias.Domain = domain
				db.DB.Create(&alias)
			}
		}
	}
}
