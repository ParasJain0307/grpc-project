package validation

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	pb "github.com/ParasJain0307/grpc-project/grpc-client/api" // Update with your actual package path
)

// ValidateUserID validates a single user ID.
func ValidateUserID(userID string) error {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}
	if id <= 0 {
		return fmt.Errorf("user ID must be greater than 0")
	}
	return nil
}

// ValidateUserIDs validates a slice of user IDs.
func ValidateUserIDs(userIDs []string) error {
	for _, userID := range userIDs {
		if err := ValidateUserID(userID); err != nil {
			return fmt.Errorf("invalid user ID: %v", err)
		}
	}
	return nil
}

// ValidateSearchCriteria validates the fields and values in a SearchCriteria.
func ValidateSearchCriteria(criteria []*pb.SearchCriteria) error {
	// Trim any leading or trailing whitespace from field name and value
	for _, criteria := range criteria {
		criteria.FieldName = trimWhitespace(criteria.FieldName)
		criteria.FieldValue = trimWhitespace(criteria.FieldValue)

		// Validate if field name is empty
		if criteria.FieldName == "" {
			return fmt.Errorf("field name cannot be empty")
		}

		// Validate if field value is empty
		if criteria.FieldValue == "" {
			return fmt.Errorf("field value cannot be empty")
		}

		// Additional custom validation rules based on field name (example rules)
		switch criteria.FieldName {
		case "fname":
			// Example: Validate first name should be alphabetic
			if !isValidAlphabetic(criteria.FieldValue) {
				return fmt.Errorf("invalid first name format: %s", criteria.FieldValue)
			}
		case "city":
			// Example: Validate city should be alphabetic with spaces
			if !isValidAlphaWithSpaces(criteria.FieldValue) {
				return fmt.Errorf("invalid city format: %s", criteria.FieldValue)
			}
		case "phone":
			// Example: Validate phone should be numeric
			if !isValidNumeric(criteria.FieldValue) {
				return fmt.Errorf("invalid phone number format: %s", criteria.FieldValue)
			}
		case "height":
			// Example: Validate height should be a float
			if !isValidFloat(criteria.FieldValue) {
				return fmt.Errorf("invalid height format: %s", criteria.FieldValue)
			}
		case "married":
			// Example: Validate married should be a boolean
			if !isValidBoolean(criteria.FieldValue) {
				return fmt.Errorf("invalid married format: %s", criteria.FieldValue)
			}
		default:
			return fmt.Errorf("unsupported field name: %s", criteria.FieldName)
		}
	}

	return nil
}

// Helper function to trim leading and trailing whitespace
func trimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

// Helper functions for specific format validations (example implementations)

func isValidAlphabetic(value string) bool {
	return isAlpha(value)
}

func isValidAlphaWithSpaces(value string) bool {
	return isAlphaWithSpaces(value)
}

func isValidNumeric(value string) bool {
	return isNumeric(value)
}

func isValidFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func isValidBoolean(value string) bool {
	_, err := strconv.ParseBool(value)
	return err == nil
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isAlphaWithSpaces(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
