package redirect

import "net/http"

func Redirect(redirectUrl string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
	}
}
