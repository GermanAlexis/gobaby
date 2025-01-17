package models

type Routes struct {
	MAIN          string
	LOGS          string
	CLOCK         string
	RESTART_CYCLE string
	LOGS_TABLE    string
	ERROR         string
	CLEAN_ERROR   string
}

var InstanceRoutes = Routes{
	MAIN:          "/main",
	LOGS:          "/logs",
	CLOCK:         "/clock",
	RESTART_CYCLE: "/restart-log",
	LOGS_TABLE:    "/logs=table",
	ERROR:         "error",
	CLEAN_ERROR:   "/clean-error",
}
