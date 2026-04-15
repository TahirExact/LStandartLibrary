package models

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Student struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
	Classes      []*Class 
}

type StudentClasses struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	ClassID   int `json:"class_id"`
}

type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
