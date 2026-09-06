package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Pnarayan-3/students-api/internal/config"
	"github.com/Pnarayan-3/students-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB

}

func New(cfg *config.Config)(*Sqlite,error){
	db,err:=sql.Open("sqlite3",cfg.StoragePath)
	if err!=nil{
		return nil,err
	}
	
	_,err=db.Exec(`CREATE TABLE IF NOT EXIST students(
	id INTEGER PRIMARY KEY AUTO-INCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err!=nil{
		return  nil,err
	}

	return &Sqlite{
		Db: db,
	},nil
}

func (s *Sqlite)CreateStudent(name string,email string,age int)(int64,error){

	stmt,err:=s.Db.Prepare("INSERT INTO students(name,email,age)VALUES(?,?,?)")
	if err!=nil{
		return 0,err
	}
	defer stmt.Close()

	result,err:=stmt.Exec(name,email,age)
	if err!=nil{
		return 0,err
	}

	lastid,err:=result.LastInsertId()
	if err!=nil{
		return 0,err
	}
	return lastid,nil

}

func (s *Sqlite)GetStudentById(id int64)(types.Student,error){
	stmt,err:=s.Db.Prepare("SELECT id,name,age FROM students WHERE id=? LIMIT 1")
	if err !=nil{
		return types.Student{},err
	}
	defer stmt.Close()
	var student types.Student

	err=stmt.QueryRow(id).Scan(&student.Id,&student.Name,&student.Email,&student.Age)
	if err!=nil{
		if err==sql.ErrNoRows{
			return types.Student{},fmt.Errorf("no student found with id %s",fmt.Sprint(id))
		}

		return types.Student{},fmt.Errorf("query error: %w",err)
	}

	return student,nil
}

func (s *Sqlite)GetStudent()([]types.Student,error){
	stmt,err:=s.Db.Prepare("SELECT id,name,email,age FROM students")
	if err!=nil{
		return nil,err
	}

	defer stmt.Close()

	rows,err:=stmt.Query()
	if err!=nil{
		return nil,err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next(){
		var student types.Student

		err:=rows.Scan(&student.Id,&student.Name,&student.Email,&student.Age)
		if err!=nil{
			return nil,err
		}
		students=append(students, student)
	}
	return students,nil
}