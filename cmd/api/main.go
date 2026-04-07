package main

import (
	"log"
	"my-project/internal/departments"
	"my-project/internal/students"
	"net/http"
)

type Server struct {
	studentHandler    *students.Handler
	departmentHandler *departments.Handler
}

func DepedencyInjection() *Server {
	//Departments domain
	departmentRepo := departments.NewMemoryRepo()
	departmentService := departments.NewService(departmentRepo)
	departmentHanler := departments.NewHandler(departmentService)

	//Students domain
	studentRepo := students.NewMemoryRepo()
	studentService := students.NewService(studentRepo)
	studentHandler := students.NewHandler(studentService)

	return &Server{
		studentHandler:    studentHandler,
		departmentHandler: departmentHanler,
	}
}

func main() {
	server := DepedencyInjection()

	mux := http.NewServeMux()

	//Students raotes
	mux.HandleFunc("GET /students/{id}", server.studentHandler.GetByID)
	mux.HandleFunc("GET /students", server.studentHandler.GetAll)
	mux.HandleFunc("POST /students", server.studentHandler.Craete)

	//Department rautes
	mux.HandleFunc("GET /departments/{id}", server.departmentHandler.GetDepartment)
	mux.HandleFunc("GET /departments", server.departmentHandler.GetDepartments)
	mux.HandleFunc("POST /departments", server.departmentHandler.Create)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
