package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidationFailed(verr error) string {
	var messages []string

	for _, err := range verr.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()

		var msg string
		switch field {
		case "Title":
			if tag == "required" {
				msg = "Title is required"
			} else if tag == "min" {
				msg = "Title must be at least 3 characters"
			} else if tag == "max" {
				msg = "Title must be at most 100 characters"
			}
		case "Description":
			if tag == "max" {
				msg = "Description must be at most 500 characters"
			}
		case "StartTime":
			if tag == "required" {
				msg = "Start time is required"
			}
		case "EndTime":
			if tag == "required" {
				msg = "End time is required"
			} else if tag == "gtfield" {
				msg = "End time must be after start time"
			}
		default:
			msg = fmt.Sprintf("Invalid value for %s", field)
		}

		messages = append(messages, msg)
	}

	finalMessage := strings.Join(messages, "; ")

	return finalMessage
}
