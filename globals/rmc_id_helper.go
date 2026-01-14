package globals

import (
	globals_rmc "github.com/PretendoNetwork/monster-hunter-4/globals/rmc"
)

func GetProtocolByID(protocolId uint16) globals_rmc.ProtocolInfoInterface {
	switch protocolId {
	case 3:
		return globals_rmc.NATTraversal{}
	case 10:
		return globals_rmc.TicketGranting{}
	case 11:
		return globals_rmc.SecureConnection{}
	case 21:
		return globals_rmc.MatchMaking{}
	case 27:
		return globals_rmc.MessageDelivery{}
	case 50:
		return globals_rmc.MatchMakingExt{}
	case 109:
		return globals_rmc.MatchmakeExtension{}
	}

	return globals_rmc.UnknownProtocol{}
}
