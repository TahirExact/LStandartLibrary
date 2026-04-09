package departments

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

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (m *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	department, err := m.service.CreteDepartment(r.Context(), input.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contedtType, applicationJson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(department)

}

func (m *Handler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	department, err := m.service.GetDepartment(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(&department)
}

func (m *Handler) GetDepartments(w http.ResponseWriter, r *http.Request) {
	departments, err := m.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(&departments)
}
