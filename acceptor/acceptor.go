package acceptor

// HighestUUID ...
var HighestUUID string

// PrepareReceive ...
func PrepareReceive(uuid string) {
	if uuid < HighestUUID {
		// Send http 400 BadRequestCode
	} else {
		HighestUUID = uuid
		// Send http 200 OK
	}
}

// AcceptReceive ...
func AcceptReceive(uuid string) {
	if uuid < HighestUUID {
		// Send http 400 BadRequestCode
	}
}
