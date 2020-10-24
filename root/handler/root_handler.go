package handler

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type RootHandler struct{}

func SetupHandler(r *mux.Router) {
	rootHandler := RootHandler{}

	r.HandleFunc("/", rootHandler.root)
}

func (h *RootHandler) root(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("root/template", "index.html")

	tmpl, err := template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "API Alkitab Indonesia",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
