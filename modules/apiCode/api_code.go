package apiCode

import "github.com/jdongdong/go-react-starter-kit/modules/errCode"

const (
	Success             int = 0
	PageNotFound            = 404
	InternalServerError     = 500
	Data_Null               = 1001
	Invalid_Params          = 1002
	Invalid_Token           = 1003
	Invalid_User            = 2001
	Server_Error            = 3001
	DB_Error                = 3002
)

func FormatApiCode(err error) int {
	if err == nil {
		return Success
	}
	_code := Server_Error
	if errCode.IsErrorDB(err) {
		_code = DB_Error
	} else if errCode.IsErrorDataNull(err) {
		_code = Data_Null
	} else if errCode.IsErrorInvalidJson(err) {
		_code = Invalid_Params
	} else if errCode.IsErrorInvalidUser(err) {
		_code = Invalid_User
	} else if errCode.IsErrorInvalidToken(err) {
		_code = Invalid_Token
	} else if errCode.IsErrorParams(err) {
		_code = Invalid_Params
	} else if errCode.IsError404(err) {
		_code = PageNotFound
	} else if errCode.IsError500(err) {
		_code = InternalServerError
	} else {
		_code = Server_Error
	}
	return _code
}
