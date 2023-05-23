package templating

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"EmailAliasManager/sharedlib/cryptography"
	"EmailAliasManager/sharedlib/db"
)

type PageData struct {
	QueryData  map[string][]string
	PostData   map[string][]string
	AuthedUser db.User
}

var FunctionMap = template.FuncMap{
	"RandomHex": cryptography.RandomHex,
	"Base64":    base64.StdEncoding.EncodeToString,
	"KeyExistsInQueryOrPostData": func(data map[string][]string, key string) bool {
		if _, ok := data[key]; ok {
			return true
		}
		return false
	},
	"Add": func(num1 int, num2 int) int {
		return num1 + num2
	},
}

func HTMLTemplate(filePath string, html fs.FS) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		parsed, err := template.New(filePath).Funcs(FunctionMap).ParseFS(html, filePath)
		if err != nil {
			fmt.Println(err)
			http.NotFound(w, r)
			return
		}

		r.ParseForm()

		var user db.User
		userContext := r.Context().Value("authedUser")
		if userContext == nil {
			user = db.User{}
		} else {
			user = userContext.(db.User)
		}

		pageData := PageData{
			QueryData:  r.URL.Query(),
			PostData:   r.Form,
			AuthedUser: user,
		}

		err = parsed.Execute(w, pageData)
		if err != nil {
			log.Println(err)
		}
	}
}
