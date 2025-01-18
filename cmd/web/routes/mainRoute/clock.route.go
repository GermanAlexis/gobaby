package mainroute

import (
	mainDomain "GoBaby/cmd/web/domain/main"
	"GoBaby/cmd/web/routes"
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
)

func ClockRender() {

	duration := mainDomain.GetDuration()
	clock := mainDomain.GetClockIntance()

	utils.SetDuration(duration)
	go utils.StartClock(clock, duration)

	routes.GetMuxIntance().HandleFunc("GET "+models.InstanceRoutes.CLOCK, mainDomain.ClockFragment)
	routes.GetMuxIntance().HandleFunc("POST "+models.InstanceRoutes.RESTART_CYCLE, mainDomain.RestartCycle)
}
