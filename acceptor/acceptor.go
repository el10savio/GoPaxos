package acceptor

import (
	"errors"
	"strconv"
)

// HighestUUID ...
var HighestUUID int64

// PrepareReceive ...
func PrepareReceive(uuid string) (bool, error) {
	if uuid == "" {
		return false, errors.New("no uuid provided")
	}

	_uuid, _ := strconv.ParseInt(uuid, 10, 64)

	if _uuid < HighestUUID {
		return false, nil
	}

	HighestUUID = _uuid
	return true, nil
}

// AcceptReceive ...
func AcceptReceive(uuid string) (bool, error) {
	if uuid == "" {
		return false, errors.New("no uuid provided")
	}

	_uuid, _ := strconv.ParseInt(uuid, 10, 64)

	return _uuid > HighestUUID, nil
}
