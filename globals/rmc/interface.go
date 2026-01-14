package globals_rmc

type ProtocolInfoInterface interface {
	Protocol() string
	GetMethodByID(uint32) string
}
