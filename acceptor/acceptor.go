package acceptor

import (
	"errors"
	"strconv"
)

// HighestUUID stores the current
// highest UUID in memory
var HighestUUID int64

// PrepareReceive checks if the incoming UUID is greater
// than the one it has ever seen before and if so
// sets the new one as the highest UUID
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

// AcceptReceive does the same as PrepareReceive
// to handle incoming Accept requests
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

	if _uuid < HighestUUID {
		return false, nil
	}

	return true, nil
}
