package routes

import "net/http"

var mux = http.NewServeMux()

func GetMuxIntance() *http.ServeMux {
	return mux
}

func GetFileServerInstance() http.Handler {
	return http.FileServer(http.Dir("ui/static"))
}
