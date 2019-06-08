package crerrors

// IFields is the interface for the Fields
type IFields interface {
	Code() string
	Level() string
	Message() string
	StatusCode() int
}

// Fields will hold the data of public data which has to be sent back to client in case of error
type fields struct {
	IFields
	statusCode    int
	publicCode    string
	publicMessage string
	debugMessage  string
}

func (f *fields) DebugMessage() string {
	return f.debugMessage
}

// StatusCode code will give the status header which has to be set for the response
func (f *fields) StatusCode() int {
	return f.statusCode
}

// Code code will give the public error code which has to be displayed to client
func (f *fields) Code() string {
	return f.publicCode
}

// Message will give the public message which has to be displayed to the client
func (f *fields) Message() string {
	return f.publicMessage
}

// Level will derive the error level of error given the status code
func (f *fields) Level() string {
	return getLevel(f.StatusCode())
}
