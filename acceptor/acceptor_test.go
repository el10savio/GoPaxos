package acceptor

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareReceive(t *testing.T) {
	HighestUUID = 123456

	expectedCondition := true
	actualCondition, actualError := PrepareReceive("123456789")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestPrepareReceive_Reject(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	actualCondition, actualError := PrepareReceive("123456")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestPrepareReceive_NoUUID(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	expectedError := errors.New("no uuid provided")
	actualCondition, actualError := PrepareReceive("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestPrepareReceive_NoHighestUUID(t *testing.T) {
	HighestUUID = 0

	expectedCondition := true
	actualCondition, actualError := PrepareReceive("123456789")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestPrepareReceive_NegativeUUID(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	expectedError := errors.New("negative uuid provided")
	actualCondition, actualError := PrepareReceive("-123456")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestPrepareReceive_InvalidUUID(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	expectedError := errors.New("invalid uuid provided")
	actualCondition, actualError := PrepareReceive("uuid")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestPrepareReceive_IsHighestUUIDUpdated(t *testing.T) {
	HighestUUID = 123456

	expectedCondition := true
	actualCondition, actualError := PrepareReceive("123456789")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)
	assert.Equal(t, int64(123456789), HighestUUID)

	HighestUUID = 0
}

func TestAcceptReceive(t *testing.T) {
	HighestUUID = 123456

	expectedCondition := true
	actualCondition, actualError := AcceptReceive("123456789")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestAcceptReceive_Reject(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	actualCondition, actualError := AcceptReceive("123456")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestAcceptReceive_NoUUID(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	expectedError := errors.New("no uuid provided")
	actualCondition, actualError := AcceptReceive("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestAcceptReceive_NoHighestUUID(t *testing.T) {
	HighestUUID = 0

	expectedCondition := true
	actualCondition, actualError := AcceptReceive("123456789")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestAcceptReceive_NegativeUUID(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	expectedError := errors.New("negative uuid provided")
	actualCondition, actualError := AcceptReceive("-123456")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}

func TestAcceptReceive_InvalidUUID(t *testing.T) {
	HighestUUID = 123456789

	expectedCondition := false
	expectedError := errors.New("invalid uuid provided")
	actualCondition, actualError := AcceptReceive("uuid")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedCondition, actualCondition)

	HighestUUID = 0
}
