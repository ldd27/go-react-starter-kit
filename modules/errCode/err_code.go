package errCode

import (
	"errors"
	"runtime"

	"github.com/jdongdong/go-lib/slog"
)

type MyError error

var (
	ErrorDB           error = errors.New("db op err")
	ErrorDataNull           = errors.New("data null")
	ErrorInvalidUser        = errors.New("invalid user")
	ErrorInvalidJson        = errors.New("invalid json")
	ErrorInvalidToken       = errors.New("invalid token")
	ErrorDataRepeat         = errors.New("data repeat")
	ErrorParams             = errors.New("invalid params")
	Error404                = errors.New("api not found")
	Error500                = errors.New("server error")
)

func CheckErrorDB(i ...interface{}) error {
	length := len(i)
	if length == 0 {
		return nil
	}
	if length == 1 {
		err, ok := i[0].(error)
		if !ok || err == nil {
			return nil
		}
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, err)
		return ErrorDB
	} else {
		if i[1] == nil {
			return nil
		}
		err, okErr := i[1].(error)
		if !okErr || err == nil {
			return nil
		}

		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, err)
		return ErrorDB
	}
}
func IsErrorDB(err error) bool {
	return err == ErrorDB
}

func CheckErrorDataNull(i ...interface{}) error {
	length := len(i)
	if length == 0 {
		return nil
	}
	if length == 1 {
		err, ok := i[0].(error)
		if !ok || err == nil {
			return nil
		}
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, err)
		return ErrorDataNull
	} else {
		if i[1] == nil {
			has, okHas := i[0].(bool)
			if !okHas || has {
				return nil
			}
			return ErrorDataNull
		} else {
			err, okErr := i[1].(error)
			if !okErr || err == nil {
				return nil
			}

			pc, file, line, ok := runtime.Caller(2)
			slog.Error(pc, file, line, ok, err)
			return ErrorDataNull
		}
	}
}
func IsErrorDataNull(err error) bool {
	return err == ErrorDataNull
}

func CheckErrorInvalidUser(err error) error {
	if err != nil {
		if IsErrorDataNull(err) {
			pc, file, line, ok := runtime.Caller(2)
			slog.Error(pc, file, line, ok, err)
			return ErrorInvalidUser
		} else {
			pc, file, line, ok := runtime.Caller(2)
			slog.Error(pc, file, line, ok, err)
			return ErrorDB
		}
	}
	return nil
}
func IsErrorInvalidUser(err error) bool {
	return err == ErrorInvalidUser
}

func CheckErrorInvalidJson(err error) error {
	if err != nil {
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, err)
		return ErrorInvalidJson
	}
	return nil
}
func IsErrorInvalidJson(err error) bool {
	return err == ErrorInvalidJson
}

func CheckErrorInvalidToken(err error) error {
	if err != nil {
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, err)
		return ErrorInvalidJson
	}
	return nil
}
func IsErrorInvalidToken(err error) bool {
	return err == ErrorInvalidToken
}

func CheckErrorDataRepeat(count int64, err error) error {
	if err != nil {
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, count, err)
		return ErrorDB
	}
	if count > 0 {
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, count, err)
		return ErrorDataRepeat
	}
	return nil
}
func IsErrorDataRepeat(count int64, err error) bool {
	return err == ErrorDataRepeat
}

func CheckErrorParams(err error) error {
	if err != nil {
		pc, file, line, ok := runtime.Caller(2)
		slog.Error(pc, file, line, ok, err)
		return ErrorParams
	}
	return nil
}
func IsErrorParams(err error) bool {
	return err == ErrorParams
}

func IsError404(err error) bool {
	return err == Error404
}

func IsError500(err error) bool {
	return err == Error500
}
