package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/oscargh945/holiday-api/domain/entities"
)

const holidaysURL = "https://api.victorsanmartin.com/feriados/en.json"

type externalHoliday struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Inalienable bool   `json:"inalienable"`
	Extra       string `json:"extra"`
}

type externalResponse struct {
	Data []externalHoliday `json:"data"`
}

type HolidayAPIClient struct {
	httpClient *http.Client
}

func NewHolidayAPIClient() *HolidayAPIClient {
	return &HolidayAPIClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *HolidayAPIClient) FetchHolidays() ([]entities.Holiday, error) {
	req, err := http.NewRequest(http.MethodGet, holidaysURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; HolidayAPI/1.0)")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch holidays: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected upstream status code: %d", resp.StatusCode)
	}

	var external externalResponse
	if err := json.NewDecoder(resp.Body).Decode(&external); err != nil {
		return nil, fmt.Errorf("failed to decode upstream response: %w", err)
	}

	holidays := make([]entities.Holiday, 0, len(external.Data))

	for _, item := range external.Data {
		holidays = append(holidays, entities.Holiday{
			Date:        item.Date,
			Title:       item.Title,
			Phone:       "",
			Type:        item.Type,
			Inalienable: item.Inalienable,
			Extra:       item.Extra,
		})
	}

	return holidays, nil
}
