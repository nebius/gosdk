package configerror

type Error struct {
	err error
}

func New(err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		err: err,
	}
}

func (e *Error) Error() string {
	return "config: " + e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}
