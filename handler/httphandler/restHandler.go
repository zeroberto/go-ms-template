package httphandler

import (
	"net/http"
)

// RestHTTPHandler is responsible for providing routines with HTTP1.1 methods
type RestHTTPHandler struct {
}

func (restHandler *RestHTTPHandler) handle(writer http.ResponseWriter, request *http.Request) error {
	// switch request.Method {
	// case http.MethodGet:

	// }
	return nil
}
