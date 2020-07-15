package acceptor

import "errors"

// HighestUUID ...
var HighestUUID string

// PrepareReceive ...
func PrepareReceive(uuid string) (bool, error) {
	if uuid == "" {
		return false, errors.New("no uuid provided")
	}

	if uuid < HighestUUID {
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
	return uuid < HighestUUID, nil
}
