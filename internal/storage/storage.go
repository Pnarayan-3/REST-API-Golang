package storage

import "github.com/Pnarayan-3/students-api/internal/types"

type Storage interface{
	CreateStudent(name string,email string,age int)(int64,error)
	GetStudentById(id int64)(types.Student,error) 
	GetStudent()([]types.Student,error)
}

