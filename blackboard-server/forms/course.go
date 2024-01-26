package forms

import(
	"github.com/go-playground/validator/v10"
)

type CourseForm struct{}

type CreateCourseForm struct {
		Coursename   string `json:"coursename" binding:"required"`
		Coursenumber string `json:"coursenumber" binding:"required,min=7,max=7"`
}

// Custom validation error messages for CreateCourseForm
func (f CourseForm) Create(form CreateCourseForm) string {
	validate := validator.New()
	err := validate.Struct(form)

	if err == nil {
		return ""
	}

	switch err := err.(type) {
	case validator.ValidationErrors:

		for _, err := range err {
			switch err.Field() {
			case "Coursename":
				return f.Coursename(err.Tag())
			case "Coursenumber":
				return f.Coursenumber(err.Tag())
			}

		}
	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

// Custom validation error messages for CreateCourseForm
func (f CourseForm) Coursename(tag string) string {
	switch tag {
	case "required":
		return "Course name is required"
	}
	return "Invalid request"
}

// Custom validation error messages for CreateCourseForm
func (f CourseForm) Coursenumber(tag string) string {
	switch tag {
	case "required":
		return "Course number is required"
	case "min":
		return "Course number must be 7 characters"
	case "max":
		return "Course number must be 7 characters"
	}
	return "Invalid request"
}
