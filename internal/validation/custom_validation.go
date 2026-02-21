package validation

import (
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidation(v *validator.Validate) {
	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	v.RegisterValidation("minInt", func(fl validator.FieldLevel) bool {
		min, _ := strconv.ParseInt(fl.Param(), 10, 64)
		if fl.Field().Int() < min{
			return false
		}
		
		return true
	})

	v.RegisterValidation("maxInt", func(fl validator.FieldLevel) bool {
		max, _ := strconv.ParseInt(fl.Param(), 10, 64)
		if fl.Field().Int() > max {
			return false
		}
		return true
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		fileHeader, ok := fl.Field().Interface().(multipart.FileHeader)
		if !ok {
			return true
		}

		fileName := fileHeader.Filename

		ext := strings.ToLower(filepath.Ext(fileName)) // .png, .jpg, ...
		if ext == "" {
			return false
		}

		ext = strings.TrimPrefix(ext, ".")

		//"png|jpg|jpeg"
		allowed := strings.Split(fl.Param(), "|")

		for _, v := range allowed {
			if strings.EqualFold(ext, strings.TrimSpace(v)) {
				return true
			}
		}

		return false
	})

	v.RegisterValidation("maxfile", func(fl validator.FieldLevel) bool {
		field := fl.Field().Interface()

		param := fl.Param()
		maxKB := 200 // default
		if param != "" {
			if n, err := strconv.Atoi(param); err == nil {
				maxKB = n
			}
		}
		maxBytes := int64(maxKB) * 1024

		switch v := field.(type) {
		// case 1: *multipart.FileHeader
		case *multipart.FileHeader:
			if v == nil {
				return true
			}
			return v.Size <= maxBytes

		// case 2: multipart.FileHeader (value)
		case multipart.FileHeader:
			return v.Size <= maxBytes

		// case 3: []*multipart.FileHeader
		case []*multipart.FileHeader:
			for _, fh := range v {
				if fh != nil && fh.Size > maxBytes {
					return false
				}
			}
			return true

		// case 4: []multipart.FileHeader (value slice)
		case []multipart.FileHeader:
			for _, fh := range v {
				if fh.Size > maxBytes {
					return false
				}
			}
			return true
		}
		return true
	})
}