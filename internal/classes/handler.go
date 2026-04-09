package classes

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const contedtType = "Content-Type"
const applicationJson = "application/json"

type Handler struct {
	service *Service
}

func NewHanler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (m *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	class, err := m.service.CreateClass(r.Context(), input.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contedtType, applicationJson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

func (m *Handler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := m.service.GetAllClasses(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(classes)
}

func (m *Handler) GetClassByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad input", http.StatusBadRequest)
		return
	}
	class, err := m.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(class)
}
