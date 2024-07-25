package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest, //400
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}
func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized, //401
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}
func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(root, msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}
func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "some thing went wrong with DB", err.Error(), "DB_ERROR")
}
func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQUEST")
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "some thing went wrong in the server", err.Error(), "INTERNAL_ERROR")
}

func ErrCannotListEntity(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotList%s", entity),
	)

}
func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotCreate%s", entity),
	)
}
func ErrCannotUpdateEntity(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotUpdate%s", entity),
	)
}
func ErrCannotDeleteEntity(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotDelete%s", entity),
	)
}
func ErrEntityExisted(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s is existed", strings.ToLower(entity)),
		fmt.Sprintf("Error%sExisted", entity),
	)
}
func ErrEntityNotFound(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Error%sNotFound", entity),
	)
}
func ErrCannotGetEntity(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotGet%s", entity),
	)
}
func ErrEntityDeleted(err error, entity string) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s is deleted", strings.ToLower(entity)),
		fmt.Sprintf("Error%sDeleted", entity),
	)
}

var RecordNotFound = errors.New("record not found")
