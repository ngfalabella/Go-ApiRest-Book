package service

import (
	"api-course/internal/model"
	"api-course/internal/store"
	"errors"
)



type Service struct {
	store store.Store
}


func New( s store.Store ) *Service {
	return &Service{
		store: s ,
	}
}


func ( s *Service ) GetAllBooks() ([] *model.Libro , error ) {
	return s.store.GetAll()
}


func( s *Service ) GetByID(id int) ( *model.Libro , error ){
	return s.store.GetByID(id)
}


func ( s *Service ) CreateBook(book model.Libro) (*model.Libro , error ) {
	if book.Title == ""{
		return nil , errors.New("necesitamos el titulo")
	}
	return s.store.CreateBook(&book)
}

func ( s *Service ) EditBook ( id int , book model.Libro ) (*model.Libro ,error) {
	if book.Title == "" {
		return nil , errors.New("necesitamos el titulo")
	}
	return s.store.UpdateBook(id,&book)
}


func ( s *Service) DeleteBook ( id int ) error {
	return s.store.DeleteBook(id)
}