package main

import (
	"log/slog"
	"os"

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
		slog.Error("failed to load holidays", "error", err)
		os.Exit(1)
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
