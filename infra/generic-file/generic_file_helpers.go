package genericfile

import (
	"encoding/json"
	"io"
	"os"

	"github.com/TudorHulban/authentication/apperrors"
)

func (store *GenericStoreFile[T]) readAll() ([]*T, error) {
	file, errRead := os.Open(store.pathFile)
	if errRead != nil {
		if os.IsNotExist(errRead) {
			return nil,
				os.WriteFile(store.pathFile, nil, 0644)
		}
		return nil, errRead
	}
	defer file.Close()

	var result []*T

	byteValue, errRead := io.ReadAll(file)
	if errRead != nil {
		return nil, errRead
	}

	if len(byteValue) == 0 {
		return nil,
			apperrors.ErrNilInput{
				InputName: store.pathFile,
			}
	}

	errUnmarshal := json.Unmarshal(byteValue, &result)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return result, nil
}

func (store *GenericStoreFile[T]) saveAll(items []*T) error {
	jsonData, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(store.pathFile, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
