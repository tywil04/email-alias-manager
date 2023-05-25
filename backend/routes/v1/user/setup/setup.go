package setup

import (
	"encoding/json"
	"net/http"

	"EmailAliasManager/lib/cryptography"
	"EmailAliasManager/lib/db"
)

type PostRequest struct {
	Token string `json:"token"`
}

type PostResponse struct {
	ID string `json:"id"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	requestData := PostRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	users := []db.User{}
	db.DB.Model(&db.User{}).Find(&users)

	if len(users) != 0 || requestData.Token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	salt, strengthenedToken := cryptography.Argon2Id(requestData.Token)

	user := db.User{
		ID:               cryptography.RandomUUID(),
		TokenSalt:        salt,
		StrengthendToken: strengthenedToken,
	}
	db.DB.Create(&user)

	response := PostResponse{
		ID: user.ID.String(),
	}
	json.NewEncoder(w).Encode(&response)
}
