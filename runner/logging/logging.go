package logging

import (
	"megacron/runner/emergencybrake"
	"megacron/system/functions"
	"megacron/system/runner"
	"runtime"
	"time"
)

type Logging struct {
	EmergencyBrake     emergencybrake.EmergencyBrakeInterface
	metrics            MetricsInterface
	commandLastRunList map[string]time.Time
}

const RUNNING_TIMEOUT_MINUTES = 5

type LoggingInterface interface {
	LogCommandsFromClosedChannel(feedBackChannel <-chan runner.RunResult)
	LogCommandsFromChannelForever(feedBackChannel <-chan runner.RunResult)
	TrackMemory(container string)
	WatchCrons()
	StartPushingMetrics()
	PushMetrics()
}

func NewLogging(em emergencybrake.EmergencyBrakeInterface, metrics MetricsInterface) LoggingInterface {
	ll := Logging{}

	if metrics == nil {
		ll.metrics = NewMetrics()
	} else {
		ll.metrics = metrics
	}

	if em == nil {
		ll.EmergencyBrake = emergencybrake.NewEmergencyBrake()
	} else {
		ll.EmergencyBrake = em
	}

	ll.commandLastRunList = map[string]time.Time{}

	return &ll
}

func (l *Logging) LogCommandsFromClosedChannel(feedBackChannel <-chan runner.RunResult) {
	for result := range feedBackChannel {
		l.logCommandResult(result)
	}
}

func (l *Logging) LogCommandsFromChannelForever(feedBackChannel <-chan runner.RunResult) {
	for {
		select {
		case result := <-feedBackChannel:
			l.logCommandResult(result)
		default:
			if l.EmergencyBrake.ShouldStop() {
				return
			}
			time.Sleep(time.Second)
		}
	}
}

func (l *Logging) logCommandResult(result runner.RunResult) {

	// passing output to stdout for kubernetes to pick up
	if result.Output != "" {
		functions.Log(result.Output)
	}

	// tracking metrics
	l.metrics.MarkCronRunningTime(result.Command, result.Elapsed)
	if result.IsSuccess {
		l.metrics.IncSuccessfulCronRun(result.Command)
	} else {
		l.metrics.IncFailedCronRun(result.Command)
	}

	// marking time when command was last run
	l.commandLastRunList[functions.FormatStringAsLabel(result.Command)] = time.Now()
}

// TrackMemory made to run in a goroutine - tracks memory every 60 secs
func (l *Logging) TrackMemory(container string) {
	var mem runtime.MemStats
	for {
		runtime.ReadMemStats(&mem)
		l.metrics.SetMemory(container, int(mem.Alloc/1024/1024))

		time.Sleep(60 * time.Second)
	}
}

// WatchCrons made to run in a goroutine
func (l *Logging) WatchCrons() {
	for {
		for command, lastRunTime := range l.commandLastRunList {
			l.metrics.ToggleRunningCron(command, lastRunTime.Add(RUNNING_TIMEOUT_MINUTES*time.Minute).After(time.Now()))
		}

		time.Sleep(60 * time.Second)
	}
}

// StartPushingMetrics pushes metrics to promethues every 60 seconds
func (l *Logging) StartPushingMetrics() {
	for {
		time.Sleep(60 * time.Second)

		l.metrics.Push()
	}
}

func (l *Logging) PushMetrics() {
	l.metrics.Push()
}
