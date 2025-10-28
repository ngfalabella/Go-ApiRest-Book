package store

import (
	"api-course/internal/model"
	"database/sql"
)

//Cual es la diferencia entre interface y struct ?

type Store interface {
	GetAll() ([]*model.Libro , error )
	GetByID( id int) (*model.Libro , error )
	CreateBook( bookCreated *model.Libro) ( *model.Libro , error )
	UpdateBook( id int ,bookUpdated *model.Libro ) ( *model.Libro , error )
	DeleteBook( id int ) error
}


//Para que es esto ?
type store struct {
	db *sql.DB
}


func New(db *sql.DB) Store{
	return &store{db :db}
}


func (s *store ) GetAll() ([]*model.Libro , error ) {
	q := `SELECT id,title,author FROM books`
	rows ,err := s.db.Query( q )

	if err != nil {
		return nil , err
	}

	defer rows.Close()

	var libros []*model.Libro
	for rows.Next(){
		var b *model.Libro
		if err := rows.Scan(&b.ID , &b.Title , &b.Author) ; err != nil {
			return nil , err
		}
		libros = append(libros ,  b)
	}
	return libros ,nil
}


func ( s *store ) GetByID( id int) (*model.Libro , error ) {
	q:= `SELECT id,title,author FROM books WHERE id=?`
	var b *model.Libro

	err := s.db.QueryRow(q,id).Scan(&b.ID ,&b.Title ,&b.Author) 
	if err != nil  {
		return nil , err
	}
	return b , nil 
}