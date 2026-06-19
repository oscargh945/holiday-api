package usecase

import (
	"testing"

	"github.com/oscargh945/holiday-api/domain/entities"
)

type fakeHolidayRepository struct {
	holidays []entities.Holiday
}

func (r *fakeHolidayRepository) FindAll() []entities.Holiday {
	return r.holidays
}

func TestListFiltersByType(t *testing.T) {
	repository := &fakeHolidayRepository{
		holidays: []entities.Holiday{
			{Date: "2024-01-01", Title: "New Year", Type: "Civil"},
			{Date: "2024-03-29", Title: "Good Friday", Type: "Religioso"},
		},
	}

	useCase := NewHolidayUseCase(repository)

	result, err := useCase.List(entities.HolidayFilter{
		Type: "Civil",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	if result[0].Type != "Civil" {
		t.Fatalf("expected Civil type, got %s", result[0].Type)
	}
}

func TestListFiltersByDateRange(t *testing.T) {
	repository := &fakeHolidayRepository{
		holidays: []entities.Holiday{
			{Date: "2024-01-01", Title: "New Year", Type: "Civil"},
			{Date: "2024-09-18", Title: "Independence Day", Type: "Civil"},
		},
	}

	useCase := NewHolidayUseCase(repository)

	result, err := useCase.List(entities.HolidayFilter{
		From: "2024-09-01",
		To:   "2024-09-30",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	if result[0].Date != "2024-09-18" {
		t.Fatalf("expected date 2024-09-18, got %s", result[0].Date)
	}
}

func TestListFiltersByTypeAndDateRange(t *testing.T) {
	repository := &fakeHolidayRepository{
		holidays: []entities.Holiday{
			{Date: "2024-01-01", Title: "New Year", Type: "Civil"},
			{Date: "2024-03-29", Title: "Good Friday", Type: "Religioso"},
			{Date: "2024-09-18", Title: "Independence Day", Type: "Civil"},
		},
	}

	useCase := NewHolidayUseCase(repository)

	result, err := useCase.List(entities.HolidayFilter{
		Type: "Civil",
		From: "2024-09-01",
		To:   "2024-09-30",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	if result[0].Date != "2024-09-18" {
		t.Fatalf("expected date 2024-09-18, got %s", result[0].Date)
	}
}

func TestListReturnsErrorForInvalidDate(t *testing.T) {
	repository := &fakeHolidayRepository{
		holidays: []entities.Holiday{
			{Date: "2024-01-01", Title: "New Year", Type: "Civil"},
		},
	}

	useCase := NewHolidayUseCase(repository)

	_, err := useCase.List(entities.HolidayFilter{
		From: "bad-date",
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
