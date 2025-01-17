package mainroute

import (
	mainDomain "GoBaby/cmd/web/domain/main"
	"GoBaby/cmd/web/routes"
	"GoBaby/internals/models"
)

func ClockRender() {
	routes.GetMuxIntance().HandleFunc("GET "+models.InstanceRoutes.CLOCK, mainDomain.ClockFragment)
}
