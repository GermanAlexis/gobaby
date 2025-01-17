package main

import (
	"GoBaby/cmd/web/routes"
	mainroute "GoBaby/cmd/web/routes/mainRoute"
	"context"
	"log"
	"log/slog"
	"net/http"
)

func InitRoutes() {
	mainroute.MainRender()
	mainroute.ClockRender()
}

func main() {
	InitRoutes()

	mux := routes.GetMuxIntance()

	fileServer := routes.GetFileServerInstance()

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	slog.Log(context.TODO(), slog.LevelInfo, "Runnnig in port 4200")
	err := http.ListenAndServe(":4200", routes.GetMuxIntance())
	log.Fatal(err)

}
