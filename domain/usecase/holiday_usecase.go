package usecase

import (
	"strings"
	"time"

	"github.com/oscargh945/holiday-api/domain/entities"
	"github.com/oscargh945/holiday-api/domain/repositories"
)

type HolidayUseCase struct {
	repository repositories.HolidayRepository
}

func NewHolidayUseCase(repository repositories.HolidayRepository) *HolidayUseCase {
	return &HolidayUseCase{
		repository: repository,
	}
}

func (u *HolidayUseCase) List(filter entities.HolidayFilter) ([]entities.Holiday, error) {
	holidays := u.repository.FindAll()

	var fromDate *time.Time
	var toDate *time.Time

	if filter.From != "" {
		parsed, err := time.Parse("2006-01-02", filter.From)
		if err != nil {
			return nil, err
		}
		fromDate = &parsed
	}

	if filter.To != "" {
		parsed, err := time.Parse("2006-01-02", filter.To)
		if err != nil {
			return nil, err
		}
		toDate = &parsed
	}

	result := make([]entities.Holiday, 0)

	for _, holiday := range holidays {
		holidayDate, err := time.Parse("2006-01-02", holiday.Date)
		if err != nil {
			continue
		}

		if filter.Type != "" && !strings.EqualFold(holiday.Type, filter.Type) {
			continue
		}

		if fromDate != nil && holidayDate.Before(*fromDate) {
			continue
		}

		if toDate != nil && holidayDate.After(*toDate) {
			continue
		}

		result = append(result, holiday)
	}

	return result, nil
}
