package httpserver

import "net/http"

var (
	msgs = map[string]string{
		"en": "Hello",
		"es": "Hola",
		"pt": "Olá",
	}
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "en"
	}

	if _, ok := msgs[lang]; !ok {
		WriteBadRequest(w, "invalid language")
		return
	}

	WriteJSON(w, http.StatusOK, HelloResponse{Message: msgs[lang]})
}
