package lead

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name		string `gorm:json"name"`
	Company		string `gorm:json"company"`
	Email		string `gorm:json"email"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.db
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.db
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func AddLead(c *fiber.Ctx) {

}

func UpdateLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.db
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}