package vo

import "reflect"

type RMap map[string]interface{}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Detail  string      `json:"detail"`
	Data    interface{} `json:"data"`
}

func (e *Result) SetCode(code int) *Result {
	e.Code = code
	return e
}

func (e *Result) SetMsg(msg string) *Result {
	e.Message = msg
	return e
}

func (e *Result) SetData(data interface{}) *Result {
	e.Data = data
	return e
}

func (e *Result) SetDetail(detail string) *Result {
	e.Detail = detail
	return e
}

func (e *Result) Json() (m RMap) {
	elem := reflect.ValueOf(e).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}

	return
}

var (
	Success = &Result{Code: 200, Message: "success"}
	Error   = &Result{Code: 1000, Message: "error"}
	Err404  = &Result{Code: 404, Message: "page not found"}
)
