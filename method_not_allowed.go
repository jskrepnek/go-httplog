package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) MethodNotAllowed(w http.ResponseWriter) {

	httpStatusCode := http.StatusMethodNotAllowed
	httpStatusName :=  StatusNameMethodNotAllowed

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}