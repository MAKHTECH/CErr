package cerr

import (
	"errors"
	"google.golang.org/grpc/codes"
)

// CErr представляет кастомную ошибку с gRPC-кодом и сообщением.
type CErr struct {
	Code    codes.Code
	Err     error
	Message string
}

// Error реализует интерфейс error.
func (e *CErr) Error() string {
	if e.Message != "" {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Err.Error()
}

// Unwrap позволяет использовать errors.Is и errors.As.
func (e *CErr) Unwrap() error {
	return e.Err
}

// Wrap позволяет добавить контекстное сообщение к ошибке.
func (e *CErr) Wrap(msg string) *CErr {
	e.Message = msg
	return e
}

// стандартная внутренняя ошибка (скрытая от пользователя)
var errInternal = &CErr{
	Err:  errors.New("internal error"),
	Code: codes.Internal,
}

//// ReturnError логирует и классифицирует ошибки.
//func ReturnError(op string, err error) error {
//	var cerr *CErr
//	if errors.As(err, &cerr) {
//		slog.Error(cerr.Message, sl.Err(cerr.Err))
//		return cerr
//	}
//
//	var pgErr *pq.Error
//	if errors.As(err, &pgErr) {
//		slog.Error("database error", sl.Err(pgErr),
//			slog.String("message", pgErr.Message),
//			slog.String("detail", pgErr.Detail),
//			slog.String("table", pgErr.Table),
//		)
//		return errInternal
//	}
//
//	slog.Error(op, sl.Err(err))
//	return errInternal
//}

// Предопределённые ошибки
var (
	ErrUserAlreadyExists = &CErr{Err: errors.New("user already exists"), Code: codes.AlreadyExists}
	ErrInvalidArgument   = &CErr{Err: errors.New("invalid argument"), Code: codes.InvalidArgument}
	ErrNotFound          = &CErr{Err: errors.New("not found"), Code: codes.NotFound}
	ErrInternal          = errInternal
)
