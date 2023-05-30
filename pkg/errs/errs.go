package errs

import "net/http"

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type MessageErrData struct {
	ErrMessage string
	ErrStatus  int
	ErrError   string
}

func (m *MessageErrData) Message() string {
	return m.ErrMessage
}

func (m *MessageErrData) Status() int {
	return m.ErrStatus
}

func (m *MessageErrData) Error() string {
	return m.ErrError
}

func NewUnauthorizedError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "NOT_AUTHORIZED",
	}
}

func NewUnathenticationError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "NOT_AUTHENTICATED",
	}
}

func NewNotFound(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "NOT_FOUND",
	}
}

func NewBadRequest(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "INTERNAL_SERVER_ERROR",
	}
}

func NewUnproccesibleEntity(message string) MessageErr {
	return &MessageErrData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "INVALID_REQUEST_BODY",
	}
}
