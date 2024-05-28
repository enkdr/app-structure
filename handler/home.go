package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type Home struct {
	templates *template.Template
}

func NewHome(templates *template.Template) *Home {
	return &Home{
		templates: templates,
	}
}

func (h *Home) HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HomePage handler")

	err := h.templates.ExecuteTemplate(w, "index.html", nil)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
