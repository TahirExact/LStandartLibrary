package models

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Student struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DepartmentID int    `json:"department_id"`
}

type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
