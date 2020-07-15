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

	_uuid, err := strconv.ParseInt(uuid, 10, 64)
	if err != nil {
		return false, errors.New("invalid uuid provided")
	}

	if _uuid < 0 {
		return false, errors.New("negative uuid provided")
	}

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

	_uuid, err := strconv.ParseInt(uuid, 10, 64)
	if err != nil {
		return false, errors.New("invalid uuid provided")
	}

	if _uuid < 0 {
		return false, errors.New("negative uuid provided")
	}

	return _uuid > HighestUUID, nil
}
