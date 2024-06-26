package genericfile

import (
	"errors"
	"fmt"
	"sync"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/helpers"
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

// TODO: change getID to return primary key
func (store *GenericStoreFile[T]) CreateItem(item *T, getID func(*T) helpers.PrimaryKey, force ...bool) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, errRead := store.readAll()
	if errRead != nil {
		return errRead
	}

	for _, reconstructedItem := range items {
		idNewItem := getID(item)

		if getID(reconstructedItem) == idNewItem {
			if len(force) == 1 && force[0] {
				return nil
			}

			return apperrors.ErrEntryAlreadyExists{
				Caller: "CreateItem",
				Entry:  item,
			}
		}
	}

	items = append(items, item)

	return store.saveAll(items)
}

func (store *GenericStoreFile[T]) UpdateItem(pk helpers.PrimaryKey, itemNew *T, getID func(*T) helpers.PrimaryKey) error {
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

func (store *GenericStoreFile[T]) DeleteItem(pk helpers.PrimaryKey, getID func(*T) helpers.PrimaryKey) error {
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

func (store *GenericStoreFile[T]) SearchItems(criterias ...func(*T) bool) ([]*T, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	items, errRead := store.readAll()
	if errRead != nil {
		return nil,
			apperrors.ErrInfrastructure{
				Issue:              errRead,
				Caller:             "SearchItems",
				NameInfrastructure: "GenericStoreFile",
			}
	}

	var result []*T

rangeItems:
	for _, item := range items {
		for _, criteria := range criterias {
			if !criteria(item) {
				continue rangeItems
			}
		}

		result = append(result, item)
	}

	if len(result) == 0 {
		return nil,
			apperrors.ErrNoEntriesFound{
				Caller: "SearchItems",
			}
	}

	return result, nil
}

func (store *GenericStoreFile[T]) SearchItem(criteriaPK func(*T) bool) (*T, error) {
	reconstructedItems, errGet := store.SearchItems(criteriaPK)
	if errGet != nil {
		return nil,
			apperrors.ErrInfrastructure{
				Issue:              errGet,
				NameInfrastructure: "GenericStoreFile",
				Caller:             "SearchItem",
			}
	}

	if len(reconstructedItems) > 1 {
		return nil,
			errors.New("duplicates found")
	}

	return reconstructedItems[0],
		nil
}
