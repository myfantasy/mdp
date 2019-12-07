package mdp

// Error type with internal
type Error struct {
	Msg           string `json:"msg,omitempty"`
	InternalError string `json:"ie,omitempty"`
}

// Error implement error interface
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.InternalError == "" {
		return e.Msg
	}
	return e.Msg + "\n" + e.InternalError
}

// ErrorS make Error from string
func ErrorS(err string) *Error {
	return &Error{Msg: err}
}

// ErrorE make Error from any error
func ErrorE(err error) *Error {
	return &Error{Msg: err.Error()}
}

// ErrorNew - Create new Error from msg and another error
func ErrorNew(msg string, internalError error) *Error {
	return &Error{Msg: msg, InternalError: internalError.Error()}
}

// ErrorNew2 - Create new Error
func ErrorNew2(msg string, internalError error, internal2Error error) *Error {
	return ErrorNew(msg, ErrorNew(internalError.Error(), internal2Error))
}
