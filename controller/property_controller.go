package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"kirmac-site-backend/services"
	"kirmac-site-backend/services/model"
	"net/http"
	"strconv"
)

type PropertyController struct {
	propertyService services.IPropertyService
}

func NewPropertyController(propertyService services.IPropertyService) *PropertyController {
	return &PropertyController{
		propertyService: propertyService,
	}
}

func (p *PropertyController) RegisterRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Get("/properties", p.getAllProperties)
	app.Get("/properties/:id", p.getPropertyById)
	app.Post("/properties", p.addProperty)
	app.Put("/properties/:id", p.updateProperty)
	app.Delete("/properties/:id", p.deleteProperty)
}

func (p *PropertyController) getAllProperties(c *fiber.Ctx) error {
	properties, err := p.propertyService.GetAllProperties()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve properties",
		})
	}
	return c.JSON(properties)
}

func (p *PropertyController) getPropertyById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	property, err := p.propertyService.GetPropertyById(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(property)
}

func (p *PropertyController) addProperty(c *fiber.Ctx) error {
	var property model.PropertyCreate
	if err := c.BodyParser(&property); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := p.propertyService.AddProperty(property)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusCreated)
}

func (p *PropertyController) updateProperty(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	var property model.PropertyCreate
	if err := c.BodyParser(&property); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	err = p.propertyService.UpdateProperty(id, property)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}

func (p *PropertyController) deleteProperty(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	deleted, err := p.propertyService.DeleteById(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"deleted": deleted})
}
