package store

import (
	"api-course/internal/model"
	"database/sql"
)

//Cual es la diferencia entre interface y struct ?


//Interface es el contrato del repositorio
type Store interface {
	GetAll() ([]*model.Libro , error )
	GetByID( id int) (*model.Libro , error )
	CreateBook( bookCreated *model.Libro) ( *model.Libro , error )
	UpdateBook( id int ,bookUpdated *model.Libro ) ( *model.Libro , error )
	DeleteBook( id int ) error
}


//Para que es esto ? implementacion
type store struct {
	db *sql.DB
}

//nos da una instancia  de nuestro store struc
func New(db *sql.DB) Store{
	return &store{db :db}
}


func (s *store ) GetAll() ([]*model.Libro , error ) {
	q := `SELECT id,title,author FROM books`
	rows ,err := s.db.Query( q )

	if err != nil {
		return nil , err
	}

	//cierra la conexion al final , ejecuta al final de la logica, patron propio del paquete sql de go
	defer rows.Close()

	var libros []*model.Libro

	//Itera sobre resultados
	for rows.Next(){
		var b model.Libro
		if err := rows.Scan(&b.ID , &b.Title , &b.Author) ; err != nil {
			return nil , err
		}
		libros = append(libros , &b)
	}
	return libros ,nil
}


func ( s *store ) GetByID( id int) (*model.Libro , error ) {
	q:= `SELECT id,title,author FROM books WHERE id=?`
	var b model.Libro

	err := s.db.QueryRow(q,id).Scan(&b.ID ,&b.Title ,&b.Author) 
	if err != nil  {
		return nil , err
	}
	return &b , nil 
}


func  ( s *store ) CreateBook ( libro *model.Libro ) ( *model.Libro , error ){
	q := `INSERT INTO TABLE books (title ,author) values (?,?) `
	resp , err := s.db.Exec(q ,  libro.Title , libro.Author )

	if err != nil {
		return nil,err
	}

	id , err  :=  resp.LastInsertId()

	if err != nil {
		return nil , err
	}


	libro.ID = int(id)
	
	return libro , nil
}


func ( s *store ) UpdateBook( id int , libro *model.Libro) ( *model.Libro , error ) {
	q := `UPDATE books SET title = ? , author = ? WHERE id`
	_, err := s.db.Exec(q,libro.Author , libro.Author , id)

	if err != nil {
		return nil , err
	}

	libro.ID = id

	return libro , nil
}


func ( s * store ) DeleteBook ( id int ) error {
	q := `DELETE from books WHERE id = ?`

	_ ,  err:=  s.db.Exec(q,id )

	if err !=  nil {
		return  err
	}

	return nil
}