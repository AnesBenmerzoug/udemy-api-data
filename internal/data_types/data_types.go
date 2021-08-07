package internal

type CourseAPIResponse struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Course `json:"results"`
}

type Course struct {
	Class  string `json:"_class"`
	Title  string `json:"title"`
	Id     int    `json:"id"`
	Url    string `json:"url"`
	IsPaid bool   `json:"is_paid"`
}
