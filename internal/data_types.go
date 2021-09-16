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
	Id                                 int          `json:"id" yaml:"course_id"`
	Url                                string       `json:"url" yaml:"url"`
	Title                              string       `json:"title" yaml:"title"`
	PublishedTitle                     string       `json:"published_title" yaml:"published_title"`
	Headline                           string       `json:"headline" yaml:"headline"`
	InstructorName                     string       `json:"instructor_name" yaml:"instructor_name"`
	Description                        *Description `json:"description" yaml:"description"`
	InstructionalLevel                 string       `json:"instructional_level" yaml:"instructional_level"`
	InstructionalLevelSimple           string       `json:"instructional_level_simple" yaml:"instructional_level_simple"`
	Prerequisites                      []string     `json:"prerequisites" yaml:"prerequisites"`
	Objectives                         []string
	PrimaryCategory                    *CourseCategory    `json:"primary_category" yaml:"primary_category"`
	PrimarySubcategory                 *CourseSubCategory `json:"primary_subcategory" yaml:"primary_subcategory"`
	HasClosedCaption                   bool               `json:"has_closed_caption" yaml:"has_closed_caption"`
	CaptionLanguages                   []string           `json:"caption_languages" yaml:"caption_languages"`
	NumLectures                        int                `json:"num_lectures" yaml:"num_lectures"`
	NumAssignments                     int                `json:"num_assignments" yaml:"num_assignments"`
	NumCodingExercises                 int                `json:"num_coding_exercises" yaml:"num_coding_exercises"`
	NumQuizzes                         int                `json:"num_quizzes" yaml:"num_quizzes"`
	PublishedTime                      string             `json:"published_time" yaml:"published_time"`
	LastUpdateDate                     string             `json:"last_update_date" yaml:"last_update_date"`
	IsPaid                             bool               `json:"is_paid" yaml:"is_paid"`
	Price                              string             `json:"price" yaml:"price"`
	PriceDetail                        CoursePriceDetail  `json:"price_detail" yaml:"price_detail"`
	NumSubscribers                     int                `json:"num_subscribers" yaml:"num_subscribers"`
	NumReviews                         int                `json:"num_reviews" yaml:"num_reviews"`
	Rating                             float32            `json:"rating" yaml:"rating"`
	AverageRating                      float32            `json:"avg_rating" yaml:"avg_rating"`
	CompletionRatio                    float32            `json:"completion_ratio" yaml:"completion_ratio"`
	ContentInfo                        string             `json:"content_info" yaml:"content_info"`
	ContentInfoShort                   string             `json:"content_info_short" yaml:"content_info_short"`
	ContentLengthVideo                 int                `json:"content_length_video" yaml:"content_length_video"`
	ContentLengthPracticeTestQuestions int                `json:"content_length_practice_test_questions" yaml:"content_length_practice_test_questions"`
	EstimatedContentLength             int                `json:"estimated_content_length" yaml:"estimated_content_length"`
	QualityStatus                      string             `json:"quality_status" yaml:"quality_status"`
	StatusLabel                        string             `json:"status_label" yaml:"status_label"`
	Image125H                          string             `json:"image_125_H" yaml:"image_125_H"`
	Image100x100                       string             `json:"image_100x100" yaml:"image_100x100"`
}

type Description string

func (description *Description) UnmarshalJSON(data []byte) error {
	escapedDescription := html.EscapeString(string(data))
	*description = Description(escapedDescription)
	return nil
}

type CoursePriceDetail struct {
	PriceString    string  `json:"price_string" yaml:"price_string"`
	Amount         float32 `json:"amount" yaml:"amount"`
	Currency       string  `json:"currency" yaml:"currency"`
	CurrencySymbol string  `json:"currency_symbol" yaml:"currency_symbol"`
}

type CourseCategory struct {
	Id           int    `json:"id" yaml:"category_id"`
	ChannelId    int    `json:"channel_id" yaml:"category_channel_id"`
	IconClass    string `json:"icon_class" yaml:"category_icon_class"`
	Title        string `json:"title" yaml:"category_title"`
	TitleCleaned string `json:"title_cleaned" yaml:"category_title_cleaned"`
	Type         string `json:"type" yaml:"category_type"`
	Url          string `json:"url" yaml:"category_url"`
}

type CourseSubCategory struct {
	Id           int    `json:"id" yaml:"sub_category_id"`
	ChannelId    int    `json:"channel_id" yaml:"sub_category_channel_id"`
	IconClass    string `json:"icon_class" yaml:"sub_category_icon_class"`
	Title        string `json:"title" yaml:"sub_category_title"`
	TitleCleaned string `json:"title_cleaned" yaml:"sub_category_title_cleaned"`
	Type         string `json:"type" yaml:"sub_category_type"`
	Url          string `json:"url" yaml:"sub_category_url"`
}

type ReviewAPIResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Reviews  []*Review `json:"results"`
}

type Review struct {
	CourseId int     `json:"course_id" yaml:"course_id"`
	Id       int     `json:"id" yaml:"review_id"`
	Title    string  `json:"title" yaml:"review_title"`
	Content  string  `json:"content" yaml:"review_content"`
	Rating   float32 `json:"rating" yaml:"review_rating"`
	Created  string  `json:"created" yaml:"review_created"`
	Modified string  `json:"modified" yaml:"review_modified"`
}
