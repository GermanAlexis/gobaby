package mainDomain

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"net/http"
)

var duration = 14400
var clockInstance = utils.NewClock()

func GetClockIntance() *utils.Clock {
	return clockInstance
}

func GetDuration() int {
	return duration
}

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckIfPath(w, r, models.InstanceRoutes.CLOCK) {
		return
	}

	files := []string{
		"ui/html/pages/main/clock.tmpl.html",
	}
	utils.ParseTemplateFiles(w, "clock", clockInstance, utils.EmptyTemplate, files...)
}

func RestartCycle(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckIfPath(w, r, models.InstanceRoutes.RESTART_CYCLE) {
		return
	}

	select {
	case <-clockInstance.Stop:
	default:
		utils.StopCountDown(clockInstance)
		clockInstance.CountDown = "04:00:00"
		utils.SetDuration(duration)
		go utils.StartClock(clockInstance, duration)
		w.Write([]byte("clock restarted"))
	}
}
