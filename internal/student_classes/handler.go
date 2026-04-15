package studentclasses

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

func (m *Handler) AssignClassToStudent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		StudentID int `json:"student_id"`
		ClassID   int `json:"class_id"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	studentClass, err := m.service.AssignClassToStudent(r.Context(), input.StudentID, input.ClassID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(contedtType, applicationJson)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(studentClass)
}

func (m *Handler) GetStudentWithClasses(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad requst", http.StatusBadRequest)
		return
	}
	student, err := m.service.GetStudentWithClasses(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(contedtType, applicationJson)
	json.NewEncoder(w).Encode(student)
}
