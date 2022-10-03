package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorEntity struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ValidationResult struct {
	Errors []*ErrorEntity
}

func (self *ValidationResult) IsPresent() bool {
	if self.Errors != nil {
		return len(self.Errors) > 0
	}

	return false
}

func (self *ValidationResult) ToJson() []map[string]interface{} {
	result := []map[string]interface{}{}

	if self.Errors == nil {
		return result
	}

	for _, err := range self.Errors {
		var entity map[string]interface{}
		buffer, err := json.Marshal(err)

		if err != nil {
			panic(err)
		}

		json.Unmarshal(buffer, &entity)

		result = append(result, entity)
	}

	return result
}

func Validate[T interface{}](params T, validate *validator.Validate) *ValidationResult {
	var errors []*ErrorEntity
	var result = ValidationResult{}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(params)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var entity ErrorEntity
			entity.Field = err.StructField()
			entity.Tag = err.Tag()
			entity.Value = fmt.Sprintf("%v", err.Value())
			err.ActualTag()
			errors = append(errors, &entity)
		}

		result.Errors = errors
		return &result
	}

	return nil
}
