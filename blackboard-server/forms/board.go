package forms

import(
	"github.com/go-playground/validator/v10"
)

type BoardForm struct{}

type CreateBoardForm struct {
	Title 	string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Custom validation error messages for CreateCourseForm
func (f BoardForm) Create(form CreateBoardForm) string {
	validate := validator.New()
	err := validate.Struct(form)

	if err == nil {
		return ""
	}

	switch err := err.(type) {
	case validator.ValidationErrors:

		for _, err := range err {
			switch err.Field() {
			case "Title":
				return f.Title(err.Tag())
			case "Desc":
				return f.Content(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Custom validation error messages for CreateCourseForm
func (f BoardForm) Title(tag string) string {
	switch tag {
	case "required":
		return "Board title is required"
	}
	return "Invalid request"
}

// Custom validation error messages for CreateCourseForm
func (f BoardForm) Content(tag string) string {
	switch tag {
	case "required":
		return "Board Content is required"
	}
	return "Invalid request"
}