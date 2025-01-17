package mainroute

import (
	mainDomain "GoBaby/cmd/web/domain/main"
	"GoBaby/cmd/web/routes"
	"GoBaby/internals/models"
)

func MainRender() {
	routes.GetMuxIntance().HandleFunc("GET "+models.InstanceRoutes.MAIN, mainDomain.MainView)
}
