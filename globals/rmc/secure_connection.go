package globals_rmc

type SecureConnection struct{}

func (protocol SecureConnection) Protocol() string {
	return "Secure Connection"
}

func (protocol SecureConnection) GetMethodByID(methodId uint32) string {
	methodTable := map[uint32]string{
		1: "Register",
		2: "RequestConnectionData",
		3: "RequestUrls",
		4: "RegisterEx",
		5: "TestConnectivity",
		6: "UpdateURLs",
		7: "ReplaceURL",
		8: "SendRequest",
	}

	method, exists := methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
