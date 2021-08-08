package internal

type CourseAPIResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Courses  []*Course `json:"results"`
}

type Course struct {
	Class                string             `json:"_class"`
	Title                string             `json:"title"`
	PublishedTitle       string             `json:"published_title"`
	Headline             string             `json:"headline"`
	Id                   int                `json:"id"`
	Url                  string             `json:"url"`
	InstructorName       string             `json:"instructor_name"`
	VisibleInstructors   []CourseInstructor `json:"visible_instructors"`
	IsPaid               bool               `json:"is_paid"`
	Price                string             `json:"price"`
	PriceDetail          CoursePriceDetail  `json:"price_detail"`
	Image125H            string             `json:"image_125_H"`
	Image240x135         string             `json:"image_240x135"`
	Image480x270         string             `json:"image_480x270"`
	IsPracticeTestCourse bool               `json:"is_practice_test_course"`
	CurriculumLectures   []string           `json:"curriculum_lectures"`
	CurriculumItems      []string           `json:"curriculum_items"`
}

type CoursePriceDetail struct {
	PriceString    string  `json:"price_string"`
	Amount         float32 `json:"amount"`
	Currency       string  `json:"currency"`
	CurrencySymbol string  `json:"currency_symbol"`
}

type CourseInstructor struct {
	Class        string `json:"_class"`
	Name         string `json:"name"`
	DisplayName  string `json:"display_name"`
	Initials     string `json:"initials"`
	Title        string `json:"title"`
	JobTitle     string `json:"job_title"`
	Image50x50   string `json:"image_50x50"`
	Image100x100 string `json:"image_100x100"`
	Url          string `json:"url"`
}
