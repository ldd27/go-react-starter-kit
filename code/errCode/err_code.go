package errCode

import (
	"IoTApi/codes/apiCode"
	"errors"
	"net/http"

	cusErr "github.com/go-errors/errors"
	"github.com/jdongdong/go-lib/slog"
	"github.com/labstack/echo"
)

var (
	errorDB = errors.New("数据库操作日程")
	//errorDataNull     = errors.New("数据为空")
	errorInvalidUser  = errors.New("无效的用户")
	errorInvalidJson  = errors.New("JSON解析错误")
	errorInvalidToken = errors.New("无效的token")
	errorDataRepeat   = errors.New("数据已存在")
	errorParams       = errors.New("非法参数")
	errorNoRecord     = errors.New("未找到符合条件数据")
	errorInternal     = echo.NewHTTPError(http.StatusInternalServerError)
)

func init() {
	cusErr.MaxStackDepth = 1
}

func NewErrorDB(err error, opts ...interface{}) error {
	return formatErr(errorDB, opts...)
}

//func NewErrorDataNull(opts ...interface{}) error {
//	return formatErr(errorDataNull, opts...)
//}

func NewErrorNoRecord(opts ...interface{}) error {
	return formatErr(errorNoRecord, opts...)
}

func NewErrorJSON(opts ...interface{}) error {
	return formatErr(errorInvalidJson, opts...)
}

func NewErrorToken(opts ...interface{}) error {
	return formatErr(errorInvalidToken, opts...)
}

func NewErrorUser(opts ...interface{}) error {
	return formatErr(errorInvalidUser, opts...)
}

func NewErrorParams(opts ...interface{}) error {
	return formatErr(errorParams, opts...)
}

func FormatApiCode(err error) int {
	if err == nil {
		return apiCode.Success
	}
	code := apiCode.ServerError
	if cusErr.Is(err, errorDB) {
		code = apiCode.DBError
		//} else if cusErr.Is(err, errorDataNull) {
		//code = apiCode.DataNull
	} else if cusErr.Is(err, errorInvalidJson) {
		code = apiCode.InvalidParams
	} else if cusErr.Is(err, errorInvalidUser) {
		code = apiCode.InvalidUser
	} else if cusErr.Is(err, errorInvalidToken) {
		code = apiCode.InvalidToken
	} else if cusErr.Is(err, errorParams) {
		code = apiCode.InvalidParams
	} else if cusErr.Is(err, echo.ErrNotFound) {
		code = apiCode.PageNotFound
	} else if cusErr.Is(err, echo.ErrMethodNotAllowed) {
		code = apiCode.MethodNotAllowed
	} else if cusErr.Is(err, errorInternal) {
		code = apiCode.InternalServerError
	} else {
		code = apiCode.ServerError
	}
	return code
}

func formatErr(formatErr error, opts ...interface{}) error {
	var err error
	var skip int
	for _, v := range opts {
		switch vv := v.(type) {
		case error: //原始错误
			err = vv
		case int:
			if vv < 0 {
				vv = 0
			}
			skip = vv
		}
	}
	newErr := cusErr.Wrap(formatErr, 2+skip)
	slog.Error("realErr:", err, "newErr:", newErr, "stack:", string(newErr.Stack()))
	return newErr
}
