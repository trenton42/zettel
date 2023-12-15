package models

type Error struct {
	id     string
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}

func (e *Error) GetType() string {
	return "error"
}

func (e *Error) GetID() string {
	if e.id == "" {
		return "UnknownError"
	}
	return e.id
}

func (e *Error) SetID(string) {}

func (e *Error) Error() string {
	return e.Title
}

func (e *Error) SetDetail(detail string) *Error {
	e.Detail = detail
	return e
}

func (e *Error) ToResponse() *Response {
	return NewResponse(e)
}

func NewError(id, message string) *Error {
	return &Error{
		id:    id,
		Title: message,
	}
}

func NewErrorResponse(id, message string) *Response {
	e := NewError(id, message)
	return NewResponse(e)
}
