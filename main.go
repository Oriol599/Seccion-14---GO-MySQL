package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Cadena de conexión a la base de datos MySQL
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Listar contactos
	handlers.ListContacts(db)

	// Obtener un contacto por ID
	contactID := 3 // ID del contacto que deseamos obtener
	handlers.GetContactByID(db, contactID)

	// Crear una instancia de Contact
	newContact := models.Contact{
		Name:  "Mireia Grass",
		Email: "mg@proton.me",
		Phone: "697-569-128",
	}

	// Registrar un nuevo contacto
	handlers.CreateContact(db, newContact)

	// Crear una instancia de Contact con los detalles actualizados
	updatedContact := models.Contact{
		Id:    4, // ID del contacto que deseas actualizar
		Name:  "Alex",
		Email: "alex@example.com",
		Phone: "987654321",
	}

	// Actualizar el contacto en la base de datos
	handlers.UpdateContact(db, updatedContact)

	// ID del contacto que deseas eliminar
	contactID = 5

	// Eliminar el contacto de la base de datos
	handlers.DeleteContact(db, contactID)

	// Listar contactos (opcional, solo para verificar que el contacto fue eliminado)
	handlers.ListContacts(db)

}
