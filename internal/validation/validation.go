package validation

import (
	"fmt"
	"strings"

	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}
	RegisterCustomValidation(v)
	return nil
}
func HandlerValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationError {
			root 	:= strings.Split(e.Namespace(), ".")[0]
			rawPath := strings.TrimPrefix(e.Namespace(), root + ".")
			parts 	:= strings.Split(rawPath, ".")
			for i, part := range parts {	
				if strings.Contains("part", "[") {
					idx 	:= strings.Index(part, "[")
					base 	:= utils.CamelToSnakeCase(part[:idx])
					index 	:= part[idx:]
					parts[i] = fmt.Sprintf("%s%s", base, index)
				} else {
					parts[i] = utils.CamelToSnakeCase(part)
				}
			}
			fieldPath := strings.Join(parts, ".")

			switch e.Tag() {
			case "uuid":
				errors[fieldPath] = fmt.Sprintf("%s is invalid uuid %s", fieldPath, e.Param())
			case "gt":
				errors[fieldPath] = fmt.Sprintf("%s must be larger %s", fieldPath, e.Param())
			case "lt":
				errors[fieldPath] = fmt.Sprintf("%s must be smaller %s", fieldPath, e.Param())
			case "slug":
				errors[fieldPath] = fmt.Sprintf("%s is invalid slug", fieldPath)
			case "required":
				errors[fieldPath] = fmt.Sprintf("%s is required", fieldPath)
			case "min":
				errors[fieldPath] = fmt.Sprintf("%s there must be at least %s character", fieldPath, e.Param())
			case "max":
				errors[fieldPath] = fmt.Sprintf("%s must not be exceeded %s character", fieldPath, e.Param())
			case "url":
				errors[fieldPath] = fmt.Sprintf("%s is invalid url", fieldPath)
			case "minInt":
				errors[fieldPath] = fmt.Sprintf("%s must be larger %s", fieldPath, e.Param())
			case "maxInt":
				errors[fieldPath] = fmt.Sprintf("%s cannot be larger %s", fieldPath, e.Param())
			case "file_ext":
				exts := strings.Split(e.Param(), " ")
				errors[fieldPath] = fmt.Sprintf("%s must have an extension of %s", fieldPath, strings.Join(exts, ", "))
			case "oneof":
				options := strings.Split(e.Param(), " ")
				errors[fieldPath] = fmt.Sprintf("%s must be one of the values: %s", fieldPath, strings.Join(options, ", "))
			case "email":
				errors[fieldPath] = fmt.Sprintf("%s must be in the correct format %s", fieldPath, fieldPath)
			case "maxfile":
				errors[fieldPath] = fmt.Sprintf("%s must be small %sKB", fieldPath, e.Param())
			}
			
		}
		return gin.H{"errors": errors}
	}
	return gin.H{
		"error": "Validation failed",
		"details": err.Error(),
	}
}