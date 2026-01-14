package globals_rmc

type MatchMakingExt struct{}

func (protocol MatchMakingExt) Protocol() string {
	return "Match Making Ext"
}

func (protocol MatchMakingExt) GetMethodByID(methodId uint32) string {
	methodTable := map[uint32]string{
		1: "EndParticipation",
		2: "GetParticipants",
		3: "GetDetailedParticipants",
		4: "GetParticipantsURLs",
		5: "GetGatheringRelations",
		6: "DeleteFromDeletions",
	}

	method, exists := methodTable[methodId]
	if exists {
		return method
	}

	return "Unknown"
}
