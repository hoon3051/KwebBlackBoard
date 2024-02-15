package forms

import(
	"github.com/go-playground/validator/v10"
	"encoding/json"
)

type UserForm struct{}

// LoginForm은 로그인 정보를 담는 폼입니다.
type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"` 
	Password string `form:"password" json:"password" binding:"required,min=4,max=30"`
}

// RegisterForm은 회원가입 정보를 담는 폼입니다.
type RegisterForm struct {
		Username      string `form:"username" json:"username" binding:"required"`
		Password      string `form:"password" json:"password" binding:"required,min=4,max=30"`
		Displayname   string `form:"displayname" json:"displayname" binding:"required,min=2,max=30"`
		Studentnumber string `form:"studentnumber" json:"studentnumber" binding:"required,min=10,max=10"`
		Isprofessor   bool 	 `form:"isprofessor" json:"isprofessor"`
}

type LoginResponse struct {
    Username string `json:"username"`
    Displayname string `json:"displayname"`
    Studentnumber string `json:"studentnumber"`
    Isprofessor bool `json:"isprofessor"`
}

// Custom validation error messages for RegisterForm
func (f UserForm) Register(form RegisterForm) string {
	validate := validator.New()
	err := validate.Struct(form)

	if err == nil {
		return ""
	}

    switch err.(type) {
    case validator.ValidationErrors:

        if _, ok := err.(*json.UnmarshalTypeError); ok {
            return "Something went wrong, please try again later"
        }

        for _, err := range err.(validator.ValidationErrors) {
            switch err.Field() {
            case "Username":
                return f.Username(err.Tag())
            case "Password":
                return f.Password(err.Tag())
            case "Displayname":
                return f.Displayname(err.Tag())
            case "Studentnumber":
                return f.Studentnumber(err.Tag())
			case "Isprofessor":
				return f.Isprofessor(err.Tag())
            }

        }
    default:
        return "Invalid request"
    }

    return "Something went wrong, please try again later"
}

// Custom validation error messages for LoginForm
func (f UserForm) Login(form LoginForm) string {
	validate := validator.New()
	err := validate.Struct(form)
	
	if err == nil {
		return ""
	}

    switch err.(type) {
    case validator.ValidationErrors:

        if _, ok := err.(*json.UnmarshalTypeError); ok {
            return "Something went wrong, please try again later"
        }

        for _, err := range err.(validator.ValidationErrors) {
            switch err.Field() {
            case "Username":
                return f.Username(err.Tag())
            case "Password":
                return f.Password(err.Tag())
            }
        }
    default:
        return "Invalid request"
    }

    return "Something went wrong, please try again later"
}

// Custom validation error messages for each field
func (f UserForm) Username(tag string) string {
    switch tag {
    case "required":
        return "Email is required"
    default:
        return "Invalid email"
    }
}

func (f UserForm) Password(tag string) string {
    switch tag {
    case "required":
        return "Password is required"
    case "min":
        return "Password is too short"
    case "max":
        return "Password is too long"
    default:
        return "Invalid password"
    }
}

func (f UserForm) Displayname(tag string) string {
	switch tag {
	case "required":
		return "Displayname is required"
	case "min":
		return "Displayname is too short"
	case "max":
		return "Displayname is too long"
	default:
		return "Invalid displayname"
	}
}

func (f UserForm) Studentnumber(tag string) string {
	switch tag {
	case "required":
		return "Studentnumber is required"
	case "min":
		return "Studentnumber must be 10 numbers"
	case "max":
		return "Studentnumber must be 10 numbers"
	default:
		return "Invalid studentnumber"
	}
}

func (f UserForm) Isprofessor(tag string) string {
	switch tag {
	default:
		return "Invalid isprofessor"
	}
}