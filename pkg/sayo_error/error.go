package sayoerror

import "fmt"

// return the error code and the error message
// If the error is not registered, then return the ErrInternalServer's code and message
func GetErrMsgByErr(err error) (int32, string) {
	code, ok := errorMp[err]
	if !ok {
		return errorMp[ErrInternalServer], ErrInternalServer.Error()
	}

	return code, err.Error()
}

// internal server error
var ()

// web info error
var (
	ErrInternalServer = fmt.Errorf("internal server error")
)

var errorMp map[error]int32 = map[error]int32{
	ErrInternalServer: 1000,
}
