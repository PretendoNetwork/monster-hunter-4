package globals_rmc

type MessageDelivery struct{}

func (protocol MessageDelivery) Protocol() string {
	return "Message Delivery"
}

func (protocol MessageDelivery) GetMethodByID(methodId uint32) string {
	methodTable := map[uint32]string{
		1: "DeliverMessage",
		2: "DeliverMessageMultiTarget",
	}

	method, exists := methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
