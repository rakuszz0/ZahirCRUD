package routes

import (
	"CRUD/handlers"
	"CRUD/pkg/mysql"
	"CRUD/repositories"

	"github.com/labstack/echo/v4"
)

func ContactRoutes(e *echo.Group) {
	contactRepository := repositories.RepositoryContact(mysql.DB)
	h := handlers.HandlerContact(contactRepository)

	e.GET("/contacts", h.GetContactList)
	e.GET("/contact/:id", h.GetContact)
	e.POST("/contact", h.CreateContact)
	e.PATCH("/contact/:id", h.UpdateContact)
	e.DELETE("/contact/:id", h.DeleteContact)
}
