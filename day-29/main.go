package main

import (
  "log"

  "github.com/gofiber/fiber/v2"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *fiber.Ctx) error {
  return c.JSON(&fiber.Map{
    "success": true,
    "albums": albums,
  })
}

func getAlbumByID(c *fiber.Ctx) error {
	id := c.AllParams()["id"]

	for _, a := range albums {
		if a.ID == id {
      return c.JSON(&fiber.Map{
        "success": true,
        "album": a,
      })
		}
	}
  return c.Status(404).JSON(&fiber.Map{
    "success": false,
    "error":   "Album not found",
  })
}

func main() {
  app := fiber.New()

  app.Get("/albums", getAlbums)
  app.Get("/albums/:id", getAlbumByID)

	log.Fatal(app.Listen(":8080"))
}
