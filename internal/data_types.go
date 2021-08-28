package internal

import (
	"golang.org/x/net/html"
)

type CourseAPIResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Courses  []*Course `json:"results"`
}

type Course struct {
	Id                                 int               `json:"id" csv:"id"`
	Url                                string            `json:"url" csv:"url"`
	Title                              string            `json:"title" csv:"title"`
	PublishedTitle                     string            `json:"published_title" csv:"published_title"`
	Headline                           string            `json:"headline" csv:"headline"`
	InstructorName                     string            `json:"instructor_name" csv:"instructor_name"`
	Description                        *Description      `json:"description" csv:"description"`
	PrimaryCategory                    *CourseCategory   `json:"primary_category" csv:"primary_category"`
	PrimarySubcategory                 *CourseCategory   `json:"primary_subcategory" csv:"primary_subcategory"`
	CaptionLanguages                   []string          `json:"caption_languages" csv:"caption_languages"`
	NumLectures                        int               `json:"num_lectures" csv:"num_lectures"`
	NumAssignments                     int               `json:"num_assignments" csv:"num_assignments"`
	NumCodingExercises                 int               `json:"num_coding_exercises" csv:"num_coding_exercises"`
	NumQuizzes                         int               `json:"num_quizzes" csv:"num_quizzes"`
	PublishedTime                      string            `json:"published_time" csv:"published_time"`
	LastUpdateDate                     string            `json:"last_update_date" csv:"last_update_date"`
	IsPaid                             bool              `json:"is_paid" csv:"is_paid"`
	Price                              string            `json:"price" csv:"price"`
	PriceDetail                        CoursePriceDetail `json:"price_detail" csv:"price_detail"`
	NumSubscribers                     int               `json:"num_subscribers" csv:"num_subscribers"`
	NumReviews                         int               `json:"num_reviews" csv:"num_reviews"`
	Rating                             float32           `json:"rating" csv:"rating"`
	AverageRating                      float32           `json:"avg_rating" csv:"avg_rating"`
	CompletionRatio                    float32           `json:"completion_ratio" csv:"completion_ratio"`
	ContentInfo                        string            `json:"content_info" csv:"content_info"`
	ContentInfoShort                   string            `json:"content_info_short" csv:"content_info_short"`
	ContentLengthVideo                 int               `json:"content_length_video" csv:"content_length_video"`
	ContentLengthPracticeTestQuestions int               `json:"content_length_practice_test_questions" csv:"content_length_practice_test_questions"`
	EstimatedContentLength             int               `json:"estimated_content_length" csv:"estimated_content_length"`
	Image125H                          string            `json:"image_125_H" csv:"image_125_H"`
	Image100x100                       string            `json:"image_100x100" csv:"image_100x100"`
}

type CoursePriceDetail struct {
	PriceString    string  `json:"price_string" csv:"price_string"`
	Amount         float32 `json:"amount" csv:"amount"`
	Currency       string  `json:"currency" csv:"currency"`
	CurrencySymbol string  `json:"currency_symbol" csv:"currency_symbol"`
}

type CourseCategory struct {
	Id           int    `json:"id" csv:"category_id"`
	ChannelId    int    `json:"channel_id" csv:"channel_id"`
	IconClass    string `json:"icon_class" csv:"icon_class"`
	Title        string `json:"title" csv:"category_title"`
	TitleCleaned string `json:"title_cleaned" csv:"category_title_cleaned"`
	Type         string `json:"type" csv:"category_type"`
	Url          string `json:"url" csv:"category_url"`
}

type Description string

func (description *Description) UnmarshalJSON(data []byte) error {
	escapedDescription := html.EscapeString(string(data))
	*description = Description(escapedDescription)
	return nil
}
