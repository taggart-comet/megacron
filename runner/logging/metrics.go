package logging

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/expfmt"
	"megacron/system/functions"
	"os"
	"strings"
	"time"
)

// MEMORY_USAGE - gauge
// how much memory every cron container consumes
const MEMORY_USAGE = "memory_usage"

// RUNNING_GOROUTINES - gauge
const RUNNING_GOROUTINES = "running_goroutines"

// NUMBER_OF_RUNNING_CRONS - gauge
// how many always running crons are running at the same time
// good metric to set an alarm to
const NUMBER_OF_RUNNING_CRONS = "number_of_running_crons"

// SUCCESSFUL_CRON_RUNS - counter
// how many crons finished with exit code 0, has command as an attribute
const SUCCESSFUL_CRON_RUNS = "successful_cron_runs"

// FAILED_CRON_RUNS - counter
// how many crons finished with exit code 1, has command as an attribute
const FAILED_CRON_RUNS = "failed_cron_runs"

// CRON_RUNNING_TIME - gauge
// for how long each cron is running
const CRON_RUNNING_TIME = "cron_running_time"

const CONTAINER_LABEL = "container_name"
const COMMAND_LABEL = "command"

type Metrics struct {
	registry          *prometheus.Registry
	memory            *prometheus.GaugeVec
	runningGoroutines *prometheus.GaugeVec
	runningCrons      *prometheus.GaugeVec
	successfulRuns    *prometheus.CounterVec
	failedRuns        *prometheus.CounterVec
	cronRunningTime   *prometheus.GaugeVec
}

type MetricsInterface interface {
	SetMemory(containerName string, valueInMegabytes int)
	ToggleRunningCron(command string, isUp bool)
	IncSuccessfulCronRun(command string)
	IncFailedCronRun(command string)
	MarkCronRunningTime(command string, runningTime time.Duration)
	Push()
}

func NewMetrics() MetricsInterface {
	metrics := Metrics{}
	metrics.initMetrics()

	return &metrics
}

func (m *Metrics) initMetrics() {

	functions.Log("Initializing metrics..")

	m.registry = prometheus.NewRegistry()

	m.memory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: MEMORY_USAGE,
		}, []string{CONTAINER_LABEL})
	m.registry.MustRegister(m.memory)

	m.runningGoroutines = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: RUNNING_GOROUTINES,
		}, []string{CONTAINER_LABEL})
	m.registry.MustRegister(m.runningGoroutines)

	m.runningCrons = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: NUMBER_OF_RUNNING_CRONS,
		}, []string{COMMAND_LABEL})
	m.registry.MustRegister(m.runningCrons)

	m.successfulRuns = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: SUCCESSFUL_CRON_RUNS,
		}, []string{COMMAND_LABEL})
	m.registry.MustRegister(m.successfulRuns)

	m.failedRuns = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: FAILED_CRON_RUNS,
		}, []string{COMMAND_LABEL})
	m.registry.MustRegister(m.failedRuns)

	m.cronRunningTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: CRON_RUNNING_TIME,
		}, []string{COMMAND_LABEL})
	m.registry.MustRegister(m.cronRunningTime)
}

func (m *Metrics) SetMemory(containerName string, valueInMegabytes int) {
	m.memory.With(prometheus.Labels{CONTAINER_LABEL: containerName}).Set(float64(valueInMegabytes))
}

func (m *Metrics) ToggleRunningCron(command string, isUp bool) {
	if isUp {
		m.runningCrons.With(prometheus.Labels{
			COMMAND_LABEL: functions.FormatStringAsLabel(command),
		}).Inc()
	} else {
		m.runningCrons.With(prometheus.Labels{
			COMMAND_LABEL: functions.FormatStringAsLabel(command),
		}).Dec()
	}
}

func (m *Metrics) IncSuccessfulCronRun(command string) {
	m.successfulRuns.With(prometheus.Labels{COMMAND_LABEL: functions.FormatStringAsLabel(command)}).Add(1)
}

func (m *Metrics) IncFailedCronRun(command string) {
	m.failedRuns.With(prometheus.Labels{COMMAND_LABEL: functions.FormatStringAsLabel(command)}).Add(1)
}

func (m *Metrics) MarkCronRunningTime(command string, runningTime time.Duration) {
	m.cronRunningTime.With(prometheus.Labels{
		COMMAND_LABEL: functions.FormatStringAsLabel(command),
	}).Set(runningTime.Seconds())
}

func (m *Metrics) Push() {

	if os.Getenv("PROMETHEUS_PUSH_ENABLED") == "false" {
		return
	}

	err := push.New(
		os.Getenv("PROMETHEUS_PUSH_GATEWAY_URL"),
		os.Getenv("PROMETHEUS_PUSH_GATEWAY_JOB"),
	).Format(expfmt.FmtText).Gatherer(m.registry).Push()

	if err != nil && !strings.Contains(err.Error(), "unexpected status code 204") {
		functions.Log("Error while pushing to prometheus: " + err.Error())
	} else {
		functions.Log("Metrics pushed to prometheus push gateway [OK]")
	}
}
