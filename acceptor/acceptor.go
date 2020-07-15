package acceptor

import (
	"errors"
	"strconv"
)

// HighestUUID ...
var HighestUUID string

// PrepareReceive ...
func PrepareReceive(uuid string) (bool, error) {
	if uuid == "" {
		return false, errors.New("no uuid provided")
	}

	_uuid, _ := strconv.ParseInt(uuid, 10, 64)
	_HighestUUID, _ := strconv.ParseInt(HighestUUID, 10, 64)

	if _uuid < _HighestUUID {
		return false, nil
	}

	HighestUUID = uuid
	return true, nil
}

// AcceptReceive ...
func AcceptReceive(uuid string) (bool, error) {
	if uuid == "" {
		return false, errors.New("no uuid provided")
	}

	_uuid, _ := strconv.ParseInt(uuid, 10, 64)
	_HighestUUID, _ := strconv.ParseInt(HighestUUID, 10, 64)

	return _uuid > _HighestUUID, nil
}
