package helpers

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/TudorHulban/authentication/apperrors"
)

func getRequestBoundary(requestHeaders map[string][]string) (string, error) {
	contenType, exists := requestHeaders["Content-Type"]
	if !exists {
		return "",
			apperrors.ErrNilInput{
				InputName: `requestHeaders["Content-Type"]`,
			}
	}

	mediaType, params, err := mime.ParseMediaType(
		strings.Join(
			contenType,
			"",
		),
	)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		if boundary, ok := params["boundary"]; ok {
			return boundary, nil
		}
	}

	return "",
		fmt.Errorf("no boundary found in Content-Type")
}

func ParseMultipartForm(formData []byte, requestHeaders map[string][]string) (map[string]string, error) {
	boundary, errCr := getRequestBoundary(requestHeaders)
	if errCr != nil {
		return nil,
			apperrors.ErrValidation{
				Issue:  errCr,
				Caller: "ParseMultipartForm",
			}
	}

	mr := multipart.NewReader(
		bytes.NewReader(formData),
		boundary,
	)

	result := make(map[string]string)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		defer part.Close()

		formField := part.FormName()

		var buf bytes.Buffer

		buf.ReadFrom(part)

		formValue := buf.String()

		if len(formValue) == 0 {
			continue
		}

		result[formField] = formValue
	}

	return result,
		nil
}
