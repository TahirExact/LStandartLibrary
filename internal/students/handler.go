package students

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

func (m *Handler) Craete(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		DeptID int    `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	student, err := m.service.CreateStudent(r.Context(), input.Name, input.DeptID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contedtType, applicationJson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (m *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	student, err := m.service.GetStudent(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(student)
}

func (m *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	students, err := m.service.GetStudents(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(students)
}
