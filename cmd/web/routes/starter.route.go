package routes

import "net/http"

var mux = http.NewServeMux()

func GetMuxIntance() *http.ServeMux {
	return mux
}
