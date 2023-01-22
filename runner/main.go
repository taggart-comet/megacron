package main

import (
	"megacron/runner/emergencybrake"
	"megacron/runner/logging"
	"megacron/runner/service"
	"megacron/system/functions"
	"megacron/system/probes"
	"megacron/system/runner"
	"os"
)

func main() {
	functions.Log("Starting.. " + os.Args[1])

	emBrake := emergencybrake.NewEmergencyBrake()

	runService := service.RunService{
		Runner:         runner.Runner{},
		Logging:        logging.NewLogging(nil, nil),
		EmergencyBrake: emBrake,
	}

	go probes.Serve()

	if shouldRunForever(os.Args[1]) {
		functions.Log("Always running..")
		runService.RunForever(os.Getenv(os.Args[1]))
		emBrake.Wait()
	} else {
		functions.Log("Running once..")
		runService.RunOnce(os.Args[1], os.Getenv(os.Args[1]), true)
	}

	functions.Log("End of script..")
}

func shouldRunForever(argument string) bool {
	return os.Getenv("ALWAYS_RUNNING_FLAG") == argument
}
