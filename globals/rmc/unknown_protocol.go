package globals_rmc

type UnknownProtocol struct{}

func (protocol UnknownProtocol) Protocol() string {
	return "Unknown"
}

func (protocol UnknownProtocol) GetMethodByID(methodId uint32) string {
	methodTable := map[uint32]string{}

	method, exists := methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
