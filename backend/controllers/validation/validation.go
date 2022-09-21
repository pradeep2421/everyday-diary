package validation

import (
	"backend/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func LoginValidation(login models.Login) error {

	// email := login.Email;
	// err = validation.Validate(email,
	// 	validation.Required,       // not empty
	// 	validation.Length(5, 100), // length between 5 and 100
	// 	is.URL,                    // is a valid URL
	// )
	return validation.ValidateStruct(&login,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&login.Email, validation.Required, validation.Length(5, 50),is.Email),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&login.Password, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		// validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		// validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)

}