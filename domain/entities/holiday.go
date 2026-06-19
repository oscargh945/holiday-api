package entities

type Holiday struct {
	Date        string `json:"date" xml:"date"`
	Title       string `json:"title" xml:"title"`
	Phone       string `json:"phone" xml:"phone"`
	Type        string `json:"type" xml:"type"`
	Inalienable bool   `json:"inalienable" xml:"inalienable"`
	Extra       string `json:"extra" xml:"extra"`
}

type HolidayFilter struct {
	Type string
	From string
	To   string
}
