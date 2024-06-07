package genericfile

import (
	"errors"
	"fmt"
	"sync"
)

type GenericStoreFile[T any] struct {
	pathFile string

	mu sync.RWMutex
}

type ParamsNewGenericStoreFile struct {
	PathStoreFile string
}

func NewGenericStoreFile[T any](params *ParamsNewGenericStoreFile) *GenericStoreFile[T] {
	return &GenericStoreFile[T]{
		pathFile: params.PathStoreFile,
	}
}

func (store *GenericStoreFile[T]) CreateFirstItem(item *T) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	return store.saveAll([]*T{item})
}

func (store *GenericStoreFile[T]) CreateItem(item *T, getID func(*T) uint64, force ...bool) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, err := store.readAll()
	if err != nil {
		return err
	}

	for _, reconstructedItem := range items {
		idNewItem := getID(item)

		if idNewItem == getID(reconstructedItem) {
			if len(force) == 1 && force[0] {
				return nil
			}

			return fmt.Errorf("item with PK: %d already exists", idNewItem)
		}
	}

	items = append(items, item)

	return store.saveAll(items)
}

func (store *GenericStoreFile[T]) UpdateItem(pk uint64, itemNew *T, getID func(*T) uint64) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, err := store.readAll()
	if err != nil {
		return err
	}

	for i, item := range items {
		if getID(item) == pk {
			items[i] = itemNew

			return store.saveAll(items)
		}
	}

	return fmt.Errorf("item with ID %d not found", pk)
}

func (store *GenericStoreFile[T]) DeleteItem(pk uint64, getID func(*T) uint64) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, err := store.readAll()
	if err != nil {
		return err
	}

	for i, item := range items {
		if getID(item) == pk {
			items = append(items[:i], items[i+1:]...)

			return store.saveAll(items)
		}
	}

	return fmt.Errorf("item with ID %d not found", pk)
}

func (store *GenericStoreFile[T]) SearchItems(criteria func(*T) bool) ([]*T, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, err := store.readAll()
	if err != nil {
		return nil,
			err
	}

	var result []*T

	for _, item := range items {
		if criteria(item) {
			result = append(result, item)
		}
	}

	return result, nil
}

func (store *GenericStoreFile[T]) SearchItem(criteriaPK func(*T) bool) (*T, error) {
	reconstructedItems, errGet := store.SearchItems(criteriaPK)
	if errGet != nil {
		return nil, errGet
	}

	if len(reconstructedItems) > 1 {
		return nil,
			errors.New("duplicates found")
	}

	return reconstructedItems[0],
		nil
}
