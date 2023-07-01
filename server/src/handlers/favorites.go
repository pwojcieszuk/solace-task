package handlers

import (
	"github.com/gofiber/fiber/v2"

	"server/database"
	"server/models"
)

func HandleError(err error, c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func HandleNotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "not found",
	})
}

func ListFavorites(c *fiber.Ctx) error {
	favorites := []models.Favorite{}
	database.DB.Db.Find(&favorites)

	return c.Status(200).JSON(favorites)
}

func GetFavorite(c *fiber.Ctx) error {
	id := c.Params("id")
	favorite := models.Favorite{}
	database.DB.Db.Find(&favorite, id)

	return c.Status(200).JSON(favorite)
}

func CreateFavorite(c *fiber.Ctx) error {
	favorite := new(models.Favorite)

	if err := c.BodyParser(favorite); err != nil {
		return HandleError(err, c)
	}

	if result := database.DB.Db.Create(&favorite); result.Error != nil {
		return HandleError(result.Error, c)
	}

	return c.Status(200).JSON(favorite)
}

func UpdateFavorite(c *fiber.Ctx) error {
	id := c.Params("id")
	favorite := models.Favorite{}
	database.DB.Db.First(&favorite, id)

	if favorite.Title == "" {
		return HandleNotFound(c)
	}

	if err := c.BodyParser(&favorite); err != nil {
		return HandleError(err, c)
	}

	database.DB.Db.Save(&favorite)

	return c.Status(200).JSON(favorite)
}

func DeleteFavorite(c *fiber.Ctx) error {
	id := c.Params("id")
	favorite := models.Favorite{}
	database.DB.Db.First(&favorite, id)

	if favorite.Title == "" {
		return HandleNotFound(c)
	}

	database.DB.Db.Unscoped().Delete(&favorite)

	return c.Status(200).JSON(fiber.Map{
		"message": "deleted",
	})
}
