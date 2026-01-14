package globals_rmc

type TicketGranting struct{}

func (protocol TicketGranting) Protocol() string {
	return "Ticket Granting"
}

func (protocol TicketGranting) GetMethodByID(methodId uint32) string {
	methodTable := map[uint32]string{
		1: "Login/ValidateAndRequestTicket",
		2: "LoginEx/ValidateAndRequestTicketWithCustomData",
		3: "RequestTicket",
		4: "GetPID",
		5: "GetName",
		6: "LoginWithContext/ValidateAndRequestTicketWithParam",
	}

	method, exists := methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
