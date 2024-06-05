package genericfile

import (
	"encoding/json"
	"io"
	"os"
)

func (store *GenericStoreFile[T]) readAll() ([]*T, error) {
	file, err := os.Open(store.pathFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil,
				os.WriteFile(store.pathFile, nil, 0644)
		}
		return nil, err
	}
	defer file.Close()

	var result []*T

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
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
