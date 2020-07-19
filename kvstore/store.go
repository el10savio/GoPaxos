package kvstore

import "errors"

// package kvstore is a simple in-memory key value
// store used to store & retrieve string entries

// Store is the implementation
// of a simple Key Value store
type Store struct {
	Map map[string]string
}

// Initialize returns an
// empty Key Value Store
func Initialize() Store {
	return Store{Map: make(map[string]string)}
}

// Set sets a key value combination to the store
func (store Store) Set(key, value string) error {
	if key == "" {
		return errors.New("empty key provided")
	}

	if value == "" {
		return errors.New("empty value provided")
	}

	store.Map[key] = value

	return nil
}

// Get obtains the value for a given key
func (store Store) Get(key string) (string, error) {
	if key == "" {
		return "", errors.New("empty key provided")
	}

	if _, ok := store.Map[key]; !ok {
		return "", errors.New("value does not exist")
	}

	return store.Map[key], nil
}

// Delete deletes a key value entry
func (store Store) Delete(key string) error {
	if key == "" {
		return errors.New("empty key provided")
	}

	if _, ok := store.Map[key]; !ok {
		return errors.New("value does not exist")
	}

	delete(store.Map, key)

	return nil
}

// Clear deletes all the key value entries
func (store Store) Clear() {
	for key := range store.Map {
		delete(store.Map, key)
	}
}
