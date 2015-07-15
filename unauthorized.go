package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Unauthorized(w http.ResponseWriter) {

	httpStatusCode := http.StatusUnauthorized
	httpStatusName :=  StatusNameUnauthorized

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}