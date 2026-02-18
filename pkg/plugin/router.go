package plugin

import (
	"log"
	"noleggio_auto/pkg/controller"
	"noleggio_auto/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(autoService service.AutoService) {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true, 
	})

	api := app.Group("/api")
	auto := api.Group("/auto")

	autoController := controller.NewAutoController(autoService)


	auto.Post("/v1", autoController.CreateAuto)         
	auto.Get("/:id/v1", autoController.GetAutoByID)    
	auto.Put("/:id/v1", autoController.UpdateAuto)     
	auto.Delete("/:id/v1", autoController.DeleteAuto)  

	
	log.Printf("Server in ascolto sulla porta :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Errore avvio Fiber: %v", err)
	}
}