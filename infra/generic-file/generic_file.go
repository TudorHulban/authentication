package genericfile

import (
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

func (store *GenericStoreFile[T]) createFirstItem(item *T) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	return store.saveAll([]*T{item})
}

func (store *GenericStoreFile[T]) CreateItem(item *T, getID func(*T) uint64) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, err := store.readAll()
	if err != nil {
		return err
	}

	for _, reconstructedItem := range items {
		idNewItem := getID(item)

		if idNewItem == getID(reconstructedItem) {
			return fmt.Errorf("item with pk: %d already exists", idNewItem)
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
