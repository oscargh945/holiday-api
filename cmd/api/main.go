package main

import (
	"log/slog"
	"os"

	"github.com/oscargh945/holiday-api/domain/entities"

	holidayclient "github.com/oscargh945/holiday-api/infrastructure/client"
	holidayrepository "github.com/oscargh945/holiday-api/infrastructure/repositories"

	"github.com/oscargh945/holiday-api/domain/usecase"
	transporthttp "github.com/oscargh945/holiday-api/transport/http"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("loading holidays from upstream service")

	apiClient := holidayclient.NewHolidayAPIClient()

	holidays, err := apiClient.FetchHolidays()
	if err != nil {
		slog.Warn("failed to load holidays from upstream services, using fallback data", "error", err)

		holidays = []entities.Holiday{
			{
				Date:        "2024-01-01",
				Title:       "Año Nuevo",
				Phone:       "",
				Type:        "Civil",
				Inalienable: true,
				Extra:       "Civil e Irrenunciable",
			},
			{
				Date:        "2024-03-29",
				Title:       "Viernes Santo",
				Phone:       "",
				Type:        "Religioso",
				Inalienable: false,
				Extra:       "Religioso",
			},
			{
				Date:        "2024-05-01",
				Title:       "Día Nacional del Trabajo",
				Phone:       "",
				Type:        "Civil",
				Inalienable: true,
				Extra:       "Civil e Irrenunciable",
			},
			{
				Date:        "2024-09-18",
				Title:       "Independencia Nacional",
				Phone:       "",
				Type:        "Civil",
				Inalienable: true,
				Extra:       "Civil e Irrenunciable",
			},
			{
				Date:        "2024-12-25",
				Title:       "Navidad",
				Phone:       "",
				Type:        "Religioso",
				Inalienable: true,
				Extra:       "Religioso e Irrenunciable",
			},
		}
	}

	slog.Info("holidays loaded successfully", "count", len(holidays))

	repository := holidayrepository.NewMemoryHolidayRepository(holidays)
	holidayUseCase := usecase.NewHolidayUseCase(repository)
	handler := transporthttp.NewHolidayHandler(holidayUseCase)
	router := transporthttp.NewRouter(handler)

	slog.Info("starting server", "port", port)

	if err := router.Run(":" + port); err != nil {
		slog.Error("server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
