package errcodes

type CodeError struct {
	code Code
	msg  string
}

// FromString parses CodeError from raw string.
//
// It ALWAYS returns CodeError object that you may
// not check if code is valid.
//
// If there is no such code, then it returns CodeError
// with CodeUnknown in it.
func FromString(raw string) *CodeError {
	code := Code(raw)
	msg, ok := codesMessages[code]
	if !ok {
		return &CodeError{
			msg:  "Unknown status",
			code: CodeUnknown,
		}
	}
	return &CodeError{
		code: code,
		msg:  msg,
	}
}

// Error implements error interface.
func (err *CodeError) Error() string {
	return err.msg
}

// Code return error's internal code.
func (err *CodeError) Code() Code {
	return err.code
}
