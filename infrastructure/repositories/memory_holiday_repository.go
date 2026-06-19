package repositories

import "github.com/oscargh945/holiday-api/domain/entities"

type MemoryHolidayRepository struct {
	holidays []entities.Holiday
}

func NewMemoryHolidayRepository(holidays []entities.Holiday) *MemoryHolidayRepository {
	return &MemoryHolidayRepository{
		holidays: holidays,
	}
}

func (r *MemoryHolidayRepository) FindAll() []entities.Holiday {
	return r.holidays
}
