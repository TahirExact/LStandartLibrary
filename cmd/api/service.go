package main

import (
	"net/http"
)

func CreteRautes(mux *http.ServeMux, server *Server) {
	mux.HandleFunc("GET /students/{id}", server.studentHandler.GetByID)
	mux.HandleFunc("GET /students", server.studentHandler.GetAll)
	mux.HandleFunc("POST /students", server.studentHandler.Craete)

	//Department rautes
	mux.HandleFunc("GET /departments/{id}", server.departmentHandler.GetDepartment)
	mux.HandleFunc("GET /departments", server.departmentHandler.GetDepartments)
	mux.HandleFunc("POST /departments", server.departmentHandler.Create)

	//Class rautes
	mux.HandleFunc("GET /classes/{id}", server.classHandler.GetClassByID)
	mux.HandleFunc("GET /classes", server.classHandler.GetAllClasses)
	mux.HandleFunc("POST /classes", server.classHandler.CreateClass)

	//StudentClass rautes
	mux.HandleFunc("POST /student/class", server.studentClassHandler.AssignClassToStudent)
	mux.HandleFunc("GET /student/class/{id}", server.studentClassHandler.GetStudentWithClasses)
}
