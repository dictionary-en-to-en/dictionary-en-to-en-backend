package inputforms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

type Search struct {
	Word string `json:"word"`
}

func (s *Search) Validation() error {
	err := validation.Errors{
		"word": validation.Validate(s.Word, validation.Required, validation.Match(regexp.MustCompile(`^[a-zA-Z\-]+$`))),
	}.Filter()

	return err
}
