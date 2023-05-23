package aliases

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"

	"EmailAliasManager/sharedlib/cryptography"
	"EmailAliasManager/sharedlib/db"
)

type PostRequest struct {
	DomainID string `json:"domainId"`
	Alias    string `json:"alias"`
	IconUrl  string `json:"iconUrl"`
}

type PostResponse struct {
	ID string `json:"id"`
}

type PatchRequest struct {
	AliasID  string `json:"aliasId"`
	Disabled bool   `json:"disabled"`
}

type PatchResponse struct {
	ID string `json:"id"`
}

type DeleteRequest struct {
	AliasID string `json:"aliasId"`
}

type DeleteResponse struct {
	ID string `json:"id"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	requestData := PostRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	if requestData.Alias == "" || requestData.DomainID == "" {
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

	alias := db.Alias{
		ID:     cryptography.RandomUUID(),
		Domain: domain,
		User:   authedUser,
		Alias:  requestData.Alias,
	}

	if requestData.IconUrl != "" {
		imageResponse, err := http.Get(requestData.IconUrl)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer imageResponse.Body.Close()

		var buffer *bytes.Buffer = new(bytes.Buffer)
		_, err = io.Copy(buffer, imageResponse.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		img, _, err := image.Decode(buffer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resizedImg := imaging.Resize(img, 64, 64, imaging.Lanczos)
		var resizedImgBuffer *bytes.Buffer = new(bytes.Buffer)

		png.Encode(resizedImgBuffer, resizedImg)

		rawImage := resizedImgBuffer.Bytes()
		mimeType := http.DetectContentType(rawImage)

		alias.IconMimeType = mimeType
		alias.IconData = rawImage
		alias.IconPresent = true
	}

	db.DB.Create(&alias)
	db.DB.Model(&authedUser).Association("Aliases").Append([]db.Alias{alias})

	response := PostResponse{
		ID: alias.ID.String(),
	}
	json.NewEncoder(w).Encode(&response)
}

func Patch(w http.ResponseWriter, r *http.Request) {
	requestData := PatchRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	if requestData.AliasID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	aliasId, err := uuid.Parse(requestData.AliasID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	alias := db.Alias{}
	db.DB.Find(&alias, db.Alias{ID: aliasId, User: authedUser})
	db.DB.Model(&alias).Update("Disabled", requestData.Disabled)

	response := PatchResponse{
		ID: alias.ID.String(),
	}
	json.NewEncoder(w).Encode(&response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	requestData := DeleteRequest{}
	json.NewDecoder(r.Body).Decode(&requestData)

	authedUser := r.Context().Value("authedUser").(db.User)

	if requestData.AliasID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	aliasId, err := uuid.Parse(requestData.AliasID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	alias := db.Alias{}
	db.DB.Delete(&alias, db.Alias{ID: aliasId, User: authedUser})

	response := DeleteResponse{
		ID: alias.ID.String(),
	}
	json.NewEncoder(w).Encode(&response)
}
