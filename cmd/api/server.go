package main

import (
	"my-project/internal/classes"
	"my-project/internal/departments"
	studentclasses "my-project/internal/student_classes"
	"my-project/internal/students"
)

type Server struct {
	studentHandler      *students.Handler
	departmentHandler   *departments.Handler
	classHandler        *classes.Handler
	studentClassHandler *studentclasses.Handler
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

	//StudentClass domain
	studentClassRepo := studentclasses.NewMemoryRepo()
	studentClassService := studentclasses.NewService(studentClassRepo, studentRepo, classRepo)
	studentClassHandler := studentclasses.NewHanler(studentClassService)

	return &Server{
		studentHandler:      studentHandler,
		departmentHandler:   departmentHanler,
		classHandler:        classHandler,
		studentClassHandler: studentClassHandler,
	}
}
