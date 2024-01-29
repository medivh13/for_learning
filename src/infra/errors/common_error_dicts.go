package errors

const (
	UNKNOWN_ERROR         ErrorCode = 0
	DATA_INVALID          ErrorCode = 4001
	FAILED_RETRIEVE_DATA  ErrorCode = 4002
	STATUS_PAGE_NOT_FOUND ErrorCode = 4003
	UNAUTHORIZED          ErrorCode = 4004
	FAILED_FORWARD_DATA   ErrorCode = 4005
	IP_ISNT_WHITELIST     ErrorCode = 4006
	RATE_LIMIT_EXCEEDED   ErrorCode = 4007
	FAILED_CREATE_DATA    ErrorCode = 4008
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	UNAUTHORIZED: {
		ClientMessage: "Unauthorized",
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_FORWARD_DATA: {
		ClientMessage: "Failed to forward data.",
		SystemMessage: "Something wrong happened while forwarding data.",
		ErrorCode:     FAILED_FORWARD_DATA,
	},
	IP_ISNT_WHITELIST: {
		ClientMessage: "Failed to forward data.",
		SystemMessage: "looks like the IP is not in whitelist.",
		ErrorCode:     IP_ISNT_WHITELIST,
	},
	RATE_LIMIT_EXCEEDED: {
		ClientMessage: "Failed to forward data.",
		SystemMessage: "Rate limit exceed.",
		ErrorCode:     RATE_LIMIT_EXCEEDED,
	},
	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
}
