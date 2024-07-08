package validation

import (
	"fmt"
	"strconv"

	pb "github.com/ParasJain0307/grpc-project/grpc-client/api"
)

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

func ValidateSearchCriteria(criterias []*pb.SearchCriteria) error {
	for _, criteria := range criterias {
		if criteria.FieldName == "" {
			return fmt.Errorf("field name cannot be empty")
		}
		if criteria.FieldValue == "" {
			return fmt.Errorf("field value cannot be empty")
		}
	}
	return nil
}
