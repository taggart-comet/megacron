package service

import (
	"megacron/runner/emergencybrake"
	"megacron/runner/logging"
	"megacron/runner/scheduler"
	"megacron/system/functions"
	"megacron/system/runner"
	"os"
	"strings"
	"sync"
	"time"
)

type RunService struct {
	Runner         runner.RunnerInterface
	Logging        logging.LoggingInterface
	EmergencyBrake emergencybrake.EmergencyBrakeInterface
}

// RunOnce runs a list of commands once
func (rs RunService) RunOnce(containerName string, commandList string, waitForAll bool) int {

	go rs.Logging.TrackMemory(containerName)
	go rs.Logging.WatchCrons()

	var waitingGroup sync.WaitGroup
	commands := strings.Split(commandList, "\n")
	feedBackChannel := make(chan runner.RunResult, len(commands))

	runCount := 0
	for _, command := range commands {
		if strings.TrimSpace(command) != "" {
			functions.Log("Starting command [" + command + "]")
			go func() {
				defer waitingGroup.Done()
				rs.Runner.Run(strings.TrimSpace(command), feedBackChannel)
			}()
			runCount++
		}
	}
	waitingGroup.Add(runCount)
	if waitForAll {
		waitingGroup.Wait()
	}
	close(feedBackChannel)

	rs.Logging.LogCommandsFromClosedChannel(feedBackChannel)

	if os.Getenv("PROMETHEUS_PUSH_ENABLED") == "true" {
		rs.Logging.PushMetrics()
	}

	return runCount
}

// RunForever runs a list of commands written in crontab format forever
func (rs RunService) RunForever(cronCommandList string) {

	go rs.Logging.TrackMemory(os.Getenv("ALWAYS_RUNNING_FLAG"))
	go rs.Logging.WatchCrons()

	if os.Getenv("PROMETHEUS_PUSH_ENABLED") == "true" {
		go rs.Logging.StartPushingMetrics()
	}

	commands := functions.SplitString(cronCommandList, "\n")
	feedBackChannel := make(chan runner.RunResult)
	for _, command := range commands {
		parsedCommand, err := scheduler.ParseScheduleFromCommand(command)
		if err != nil {
			continue
		}

		functions.Log("Launching cron [" + parsedCommand.CommandString + "] with schedule [" + parsedCommand.ScheduleString + "]")
		go rs.runBySchedule(parsedCommand, feedBackChannel)
	}

	// logging the output of our commands
	rs.Logging.LogCommandsFromChannelForever(feedBackChannel)
}

// runs a single command forever by given schedule
func (rs RunService) runBySchedule(
	scheduledCommand scheduler.CommandSchedule,
	feedBackChannel chan runner.RunResult,
) {
	rs.EmergencyBrake.Start()
	for {
		if rs.EmergencyBrake.ShouldStop() {
			rs.EmergencyBrake.Stop()
			break
		}

		nextRun := scheduledCommand.CronExpression.Next(time.Now())
		time.Sleep(nextRun.Sub(time.Now()))

		if rs.EmergencyBrake.ShouldStop() {
			rs.EmergencyBrake.Stop()
			break
		}
		rs.Runner.Run(scheduledCommand.CommandString, feedBackChannel)
	}
}
