package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup() {
	app := fiber.New()

	app.Get("/api/allnotes", GetAllNotes)
	app.Get("/api/note/:id", GetNoteById)

	app.Listen(":3000")
}
