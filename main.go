package main

import (
	"api-course/internal/service"
	"api-course/internal/store"
	"api-course/internal/transport"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	// separacion of concerns // separacion de responsabilidades
	// 3 layers

	// http 	transport
	//todo lo relacionado a rutas y recursos va al transport

	// busines logic 	service
	// va toda la logica, buscar libro, mezclar libro, con otro recurso y toda esa complejdidad va a parar aca

	// base de datos store
	//todo lo relacionado a base de datos

	//separado de nuestra estructura => model//
	//defnie estructura de datos

	//Conectar a sqllite => se escribe archivo de forma local cumpliendo la funcion de bd

	db, err := sql.Open("sqlite3", "./books.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// crear table si no existe

	q := `
CREATE TABLE IF NOT EXISTS books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	author TEXT NOT NULL
);
`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inyectar nuestras dependencias

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	//Configurar rutas

	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/book/", bookHandler.HandleBookById)

	fmt.Println("Servidor corriendo en 8080")

	//Empezar y escuchar servidor
	log.Fatal(http.ListenAndServe(":8080", nil))

}
