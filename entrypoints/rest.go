package entrypoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theleeeo/form-forge/app"
)

func NewRestHandler(app *app.App) *restHandler {
	h := &restHandler{
		app: app,
	}

	return h
}

func (h *restHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /form/{id}", h.getRenderedForm)
	mux.HandleFunc("POST /submit/{id}", h.handleSubmit)
}

type restHandler struct {
	app *app.App
}

func (h *restHandler) handleSubmit(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("error parsing form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	log.Printf("Recieved response for post: %s", id)
	for k, v := range r.PostForm {
		log.Printf("Key: %s, Value: %v", k, v)
	}

	h.app.SubmitResponse(r.Context(), id, r.PostForm)
}

func (h *restHandler) getRenderedForm(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	tpl, err := h.app.TemplateForm(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tpl)
}
