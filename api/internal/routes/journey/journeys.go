package journeyRoutes

import (
	journeyHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/journey"
	debtHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/journey/debt"
	noteHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/journey/note"
	pointHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/journey/point"
	travelerHandler "github.com/Central-University-IT-prod/backend-oleg-top/api/internal/handlers/journey/traveler"
	"github.com/gofiber/fiber/v3"
)

func SetupJourneyRoutes(router fiber.Router) {
	journeyRouter := router.Group("/journey")
	journeyRouter.Post("/create", journeyHandler.CreateJourney)
	journeyRouter.Get("/chat/:chat_id", journeyHandler.ChatJourney)
	journeyRouter.Get("/staticmap/:journey_name/:chat_id", journeyHandler.StaticMapJourney)
	journeyRouter.Get("/hotels/:journey_name", journeyHandler.HotelsJourney)
	journeyRouter.Get("/weather/:journey_name", journeyHandler.WeatherJourney)
	journeyRouter.Get("/restaurants/:journey_name", journeyHandler.RestaurantJourney)
	journeyRouter.Get("/venues/:journey_name", journeyHandler.VenueJourney)
	journeyRouter.Get("/about/:journey_name", journeyHandler.AboutJourney)
	journeyRouter.Delete("/remove/:journey_name", journeyHandler.RemoveJourney)

	pointsRouter := journeyRouter.Group("/point")
	pointsRouter.Post("/add", pointHandler.AddPoint)
	pointsRouter.Delete("/remove", pointHandler.RemovePoint)

	travelerRouter := journeyRouter.Group("/traveler")
	travelerRouter.Post("/add", travelerHandler.AddTraveler)

	noteRouter := journeyRouter.Group("/note")
	noteRouter.Post("/create", noteHandler.CreateNote)
	noteRouter.Post("/upload_file/:note_id", noteHandler.UploadFileNote)

	debtRouter := journeyRouter.Group("/debt")
	debtRouter.Post("/create", debtHandler.CreateDebt)
	debtRouter.Post("/pay", debtHandler.PayDebt)
	debtRouter.Get("/list/:chat_id", debtHandler.ListDebts)
}
