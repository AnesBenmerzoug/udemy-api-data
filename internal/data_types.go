package internal

type CourseAPIResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Courses  []*Course `json:"results"`
}

type Course struct {
	Class                string             `json:"_class" csv:"class"`
	Title                string             `json:"title" csv:"title"`
	PublishedTitle       string             `json:"published_title" csv:"published_title"`
	Headline             string             `json:"headline" csv:"headline"`
	Id                   int                `json:"id" csv:"id"`
	Url                  string             `json:"url" csv:"url"`
	InstructorName       string             `json:"instructor_name" csv:"instructor_name"`
	VisibleInstructors   []CourseInstructor `json:"visible_instructors" csv:"visible_instructors"`
	IsPaid               bool               `json:"is_paid" csv:"is_paid"`
	Price                string             `json:"price" csv:"price"`
	PriceDetail          CoursePriceDetail  `json:"price_detail" csv:"price_detail"`
	Image125H            string             `json:"image_125_H" csv:"image_125_H"`
	Image240x135         string             `json:"image_240x135" csv:"image_240x135"`
	Image480x270         string             `json:"image_480x270" csv:"image_480x270"`
	IsPracticeTestCourse bool               `json:"is_practice_test_course" csv:"is_practice_test_course"`
	CurriculumLectures   []string           `json:"curriculum_lectures" csv:"curriculum_lectures"`
	CurriculumItems      []string           `json:"curriculum_items" csv:"curriculum_items"`
}

type CoursePriceDetail struct {
	PriceString    string  `json:"price_string" csv:"price_string"`
	Amount         float32 `json:"amount" csv:"amount"`
	Currency       string  `json:"currency" csv:"currency"`
	CurrencySymbol string  `json:"currency_symbol" csv:"currency_symbol"`
}

type CourseInstructor struct {
	Class        string `json:"_class" csv:"class"`
	Name         string `json:"name" csv:"name"`
	DisplayName  string `json:"display_name" csv:"display_name"`
	Initials     string `json:"initials" csv:"initials"`
	Title        string `json:"title" csv:"title"`
	JobTitle     string `json:"job_title" csv:"job_title"`
	Image50x50   string `json:"image_50x50" csv:"image_50x50"`
	Image100x100 string `json:"image_100x100" csv:"image_100x100"`
	Url          string `json:"url" csv:"url"`
}
