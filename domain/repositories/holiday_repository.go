package repositories

import "github.com/oscargh945/holiday-api/domain/entities"

type HolidayRepository interface {
	FindAll() []entities.Holiday
}
