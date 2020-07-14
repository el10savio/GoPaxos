package kvstore

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	store Store
)

func init() {
	store = Initialize()
}

func TestGet(t *testing.T) {
	store.Set("xx", "yy")

	expectedValue := "yy"
	actualValue, actualError := store.Get("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestGet_UpdatedValue(t *testing.T) {
	store.Set("xx", "yy")
	store.Set("xx", "zz")

	expectedValue := "zz"
	actualValue, actualError := store.Get("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestGet_NoValue(t *testing.T) {
	_ = store.Set("xx", "")

	expectedValue := ""
	expectedError := errors.New("value does not exist")
	actualValue, actualError := store.Get("xx")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestGet_NoKey(t *testing.T) {
	expectedValue := ""
	expectedError := errors.New("empty key provided")
	actualValue, actualError := store.Get("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestSet_NoValue(t *testing.T) {
	expectedError := errors.New("empty value provided")
	actualError := store.Set("xx", "")

	assert.Equal(t, expectedError, actualError)

	store.Clear()
}

func TestClear(t *testing.T) {
	store.Set("xx1", "yy")
	store.Set("xx2", "yy")
	store.Clear()

	expectedValue := ""
	expectedError := errors.New("value does not exist")
	actualValue, actualError := store.Get("xx1")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestClear_EmptyStore(t *testing.T) {
	store.Clear()

	expectedValue := ""
	expectedError := errors.New("value does not exist")
	actualValue, actualError := store.Get("xx1")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestDelete(t *testing.T) {
	store.Set("xx", "yy")
	store.Delete("xx")

	expectedValue := ""
	expectedError := errors.New("value does not exist")
	actualValue, actualError := store.Get("xx")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	store.Clear()
}

func TestDelete_InvalidKey(t *testing.T) {
	store.Set("xx", "yy")

	expectedError := errors.New("value does not exist")
	actualError := store.Delete("zz")

	assert.Equal(t, expectedError, actualError)

	store.Clear()
}

func TestDelete_EmptyKey(t *testing.T) {
	store.Set("xx", "yy")

	expectedError := errors.New("empty key provided")
	actualError := store.Delete("")

	assert.Equal(t, expectedError, actualError)

	store.Clear()
}

func TestDelete_EmptyStore(t *testing.T) {
	expectedError := errors.New("empty key provided")
	actualError := store.Delete("")

	assert.Equal(t, expectedError, actualError)

	store.Clear()
}
