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
	ErrInternalServer      = fmt.Errorf("internal server error")
	ErrDuplicateIdentifier = fmt.Errorf("duplicate identifier")
	ErrRegisterFailed      = fmt.Errorf("modules register failed")
	ErrUnknownType         = fmt.Errorf("unknown register type")
	ErrInvalidRole         = fmt.Errorf("invalid register role")
)

var errorMp map[error]int32 = map[error]int32{
	ErrInternalServer:      1000,
	ErrDuplicateIdentifier: 1001,
	ErrRegisterFailed:      1002,
	ErrUnknownType:         1003,
	ErrInvalidRole:         1004,
}
