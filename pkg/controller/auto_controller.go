package controller

import (
	"noleggio_auto/pkg/dto"
	"noleggio_auto/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type AutoController struct {
	srv service.AutoService
}

func NewAutoController(srv service.AutoService) *AutoController {
	return &AutoController{srv: srv}
}


func (ac *AutoController) CreateAuto(c *fiber.Ctx) error {
	var input dto.CreateAutoRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Corpo della richiesta non valido",
		})
	}

	auto, err := ac.srv.CreaNuovaAuto(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(auto)
}

func (ac *AutoController) GetAutoByID(c *fiber.Ctx) error {
	id := c.Params("id")

	auto, err := ac.srv.OttieniAuto(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Auto non trovata",
		})
	}

	return c.JSON(auto)
}

func (ac *AutoController) UpdateAuto(c *fiber.Ctx) error {
	id := c.Params("id")
	var input dto.UpdateAutoRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dati di aggiornamento non validi",
		})
	}

	err := ac.srv.AggiornaAuto(c.Context(), id, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Errore durante l'aggiornamento: " + err.Error(),
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Auto aggiornata con successo",
	})
}


func (ac *AutoController) DeleteAuto(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ac.srv.RimuoviAuto(c.Context(), id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Impossibile eliminare l'auto",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}