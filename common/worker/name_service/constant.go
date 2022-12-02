package name_service

const (
	ErrUnregisterName     = "unregistered name"
	ErrNotEvmAddress      = "the address provided is invalid. You can use the address begin with 0x or Name Service Resolution refer to the documentation"
	ErrNotSupportNS       = "currently not supported, please refer to the documentation for supported"
	ErrNotSupportContract = "contract addresses are not currently supported"
	ErrNotResolver        = "not configure forward resolution"

	ReferDoc = "https://docs.rss3.io/docs/name-service-resolution"
)
