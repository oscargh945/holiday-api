package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/oscargh945/holiday-api/domain/entities"
)

var holidaysURLs = []string{
	"https://api.victorsanmartin.com/feriados/en.json",
	"https://api.boostr.cl/feriados/en.json",
}

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
	var lastErr error

	for _, url := range holidaysURLs {
		holidays, err := c.fetchFromURL(url)
		if err == nil {
			return holidays, nil
		}

		lastErr = err
	}

	return nil, fmt.Errorf("all upstream holiday services failed: %w", lastErr)
}

func (c *HolidayAPIClient) fetchFromURL(url string) ([]entities.Holiday, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", url, err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; HolidayAPI/1.0)")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch holidays from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected upstream status code: %d, requested_url: %s, final_url: %s",
			resp.StatusCode,
			url,
			resp.Request.URL.String(),
		)
	}

	var external externalResponse
	if err := json.NewDecoder(resp.Body).Decode(&external); err != nil {
		return nil, fmt.Errorf("failed to decode upstream response from %s: %w", url, err)
	}

	if len(external.Data) == 0 {
		return nil, fmt.Errorf("upstream service returned empty holidays data from %s", url)
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
