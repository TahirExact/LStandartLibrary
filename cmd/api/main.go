package main

import (
	"log"
	"my-project/internal/classes"
	"my-project/internal/departments"
	"my-project/internal/students"
	"net/http"
)

type Server struct {
	studentHandler    *students.Handler
	departmentHandler *departments.Handler
	classHandler      *classes.Handler
}

func DepedencyInjection() *Server {
	//Departments domain
	departmentRepo := departments.NewMemoryRepo()
	departmentService := departments.NewService(departmentRepo)
	departmentHanler := departments.NewHandler(departmentService)

	//Classes domain
	classRepo := classes.NewMemoryRepo()
	classService := classes.NewService(classRepo)
	classHandler := classes.NewHanler(classService)

	//Students domain
	studentRepo := students.NewMemoryRepo()
	studentService := students.NewService(studentRepo, departmentRepo)
	studentHandler := students.NewHandler(studentService)

	return &Server{
		studentHandler:    studentHandler,
		departmentHandler: departmentHanler,
		classHandler:      classHandler,
	}
}

func main() {
	server := DepedencyInjection()

	mux := http.NewServeMux()

	//Students rautes
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

	log.Fatal(http.ListenAndServe(":8080", mux))
}
