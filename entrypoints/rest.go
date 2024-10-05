package entrypoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
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

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("error parsing form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err := h.app.SubmitResponse(r.Context(), uid, r.PostForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("response submitted for form %s", id)
}

func (h *restHandler) getRenderedForm(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not parse id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	tpl, err := h.app.TemplateForm(r.Context(), uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(tpl); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
