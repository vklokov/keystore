package validations

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

const (
	KIND_VALIDATION = "validation"
)

type VaError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
	Kind    string `json:"kind"`
}

type Result struct {
	Errors []*VaError
}

type VaMap = map[string]interface{}

func (self *Result) HasErrors() bool {
	return len(self.Errors) > 0
}

func (self *Result) ToJson() []VaMap {
	result := []VaMap{}

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

func CreateResult(errors []*VaError) *Result {
	result := &Result{}
	result.Errors = errors
	return result
}

func Validate[T interface{}](params T, v *validator.Validate) *Result {
	errors := []*VaError{}
	result := &Result{}

	err := v.Struct(params)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			entity := &VaError{}
			entity.Field = err.StructField()
			entity.Tag = err.Tag()
			entity.Kind = KIND_VALIDATION
			errors = append(errors, entity)
		}
		result.Errors = errors
		return result
	}

	return nil
}
