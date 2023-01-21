package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"megacron/mocks"
	"megacron/runner/logging"
	"os"
	"testing"
)

func TestRunOnce(t *testing.T) {

	commands := `
echo 1
echo 2
sleep 1
echo 3
`
	runnerMock := mocks.NewRunnerInterface(t)
	loggerMock := mocks.NewLoggingInterface(t)
	runnerMock.On("Run", mock.AnythingOfType("string"), mock.Anything).Maybe()
	loggerMock.On("LogCommandsFromClosedChannel", mock.Anything)
	loggerMock.On("PushMetrics", mock.Anything)
	loggerMock.On("TrackMemory", mock.Anything).Maybe()
	loggerMock.On("WatchCrons", mock.Anything).Maybe()

	service := RunService{
		Runner:  runnerMock,
		Logging: loggerMock,
	}
	os.Setenv("PROMETHEUS_PUSH_ENABLED", "true")
	runCount := service.RunOnce("TEST_RUN", commands, false)
	assert.Equal(t, 4, runCount)
	loggerMock.AssertNumberOfCalls(t, "LogCommandsFromClosedChannel", 1)
}

func TestEmergencyBraking(t *testing.T) {
	commands := `
* * * * * echo 1
*/5 * * * * echo 2
* * * * * sleep 1
* * * * * echo 3
`
	runnerMock := mocks.NewRunnerInterface(t)
	emMock := mocks.NewEmergencyBrakeInterface(t)
	metricsMock := mocks.NewMetricsInterface(t)
	loggingInstance := logging.NewLogging(emMock, metricsMock)
	runnerMock.On("Run", mock.AnythingOfType("string"), mock.Anything).Maybe()
	emMock.On("ShouldStop").Return(true)
	metricsMock.On("SetMemory", mock.Anything, mock.Anything).Maybe()
	metricsMock.On("ToggleRunningCron", mock.Anything).Maybe()
	metricsMock.On("Push", mock.Anything).Maybe()

	os.Setenv("PROMETHEUS_PUSH_ENABLED", "true")

	service := RunService{
		Runner:         runnerMock,
		Logging:        loggingInstance,
		EmergencyBrake: emMock,
	}

	service.RunForever(commands)
}
