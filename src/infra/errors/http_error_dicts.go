package errors

import (
	"net/http"
)

var httpCode = map[ErrorCode]int{
	UNKNOWN_ERROR:            http.StatusInternalServerError,
	DATA_INVALID:             http.StatusBadRequest,
	STATUS_PAGE_NOT_FOUND:    http.StatusNotFound,
	UNAUTHORIZED:             http.StatusUnauthorized,
	FAILED_RETRIEVE_DATA:     http.StatusInternalServerError,
	IP_ISNT_WHITELIST:        http.StatusBadRequest,
	RATE_LIMIT_EXCEEDED:      http.StatusTooManyRequests,
}
