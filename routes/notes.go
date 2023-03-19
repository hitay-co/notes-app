package routes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Completed bool      `json:"completed"`
}

func GetAllNotes(c *fiber.Ctx) error {
	note := Note{
		ID:        5,
		Title:     "My first note",
		Note:      "This is a note!",
		CreatedAt: time.Time{},
		Completed: false,
	}

	return c.JSON(note)

}

func GetNoteById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		fmt.Println("error")
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	fmt.Println("id : %d\n", id)
	// SQL Query for the actual note itself.
	note := Note{
		ID:        2,
		Title:     "Another Note",
		Note:      "My second awesome note!",
		CreatedAt: time.Time{},
		Completed: true,
	}
	return c.JSON(note)
}
