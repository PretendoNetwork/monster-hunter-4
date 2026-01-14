package globals_rmc

type NATTraversal struct{}

func (protocol NATTraversal) Protocol() string {
	return "NAT Traversal"
}

func (protocol NATTraversal) GetMethodByID(methodId uint32) string {
	methodTable := map[uint32]string{
		1: "RequestProbeInitiation",
		2: "InitiateProbe",
		3: "RequestProbeInitiationExt",
		4: "ReportNATTraversalResult",
		5: "ReportNATProperties",
		6: "GetRelaySignatureKey",
		7: "ReportNATTraversalDetail",
	}

	method, exists := methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
