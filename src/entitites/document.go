package entitites

import validation "github.com/go-ozzo/ozzo-validation"

type Document struct {
	DocumentNo    string
	DocumentTitle string
	Notes         string
}

func (d Document) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.DocumentNo, validation.Required),
		validation.Field(&d.DocumentTitle, validation.Required),
	)
}
