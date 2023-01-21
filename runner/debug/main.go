package main

import (
	"megacron/runner/emergencybrake"
	"megacron/runner/logging"
	"megacron/runner/service"
	"megacron/system/runner"
	"os"
)

func main() {

	os.Setenv("PROMETHEUS_PUSH_GATEWAY_URL", "http://dev-vminsert.sberhealth.pro/insert/0/prometheus/api/v1/import/prometheus")
	os.Setenv("PROMETHEUS_PUSH_GATEWAY_JOB", "marketplace_test_cron")
	os.Setenv("PROMETHEUS_PUSH_ENABLED", "true")

	runService := service.RunService{
		Runner:         runner.Runner{},
		Logging:        logging.NewLogging(nil, nil),
		EmergencyBrake: emergencybrake.NewEmergencyBrake(),
	}

	runService.RunOnce("debug", "echo 123", true)
}
