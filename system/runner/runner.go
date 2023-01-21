package runner

import (
	"megacron/system/functions"
	"os/exec"
	"strings"
	"time"
)

type Runner struct{}

type RunnerInterface interface {
	Run(commandWithArgs string, feedBackChannel chan<- RunResult)
}

type RunResult struct {
	IsSuccess bool
	Output    string
	Elapsed   time.Duration
	Command   string
}

func (r Runner) Run(commandWithArgs string, feedBackChannel chan<- RunResult) {

	if commandWithArgs == "" {
		return
	}

	start := time.Now()
	parts := functions.SplitString(commandWithArgs, " ")
	run := exec.Command(parts[0], parts[1:]...)
	output, err := run.Output()

	feedBackChannel <- RunResult{
		IsSuccess: err == nil,
		Output:    strings.TrimSuffix(functions.BytesToString(output), "\n"),
		Elapsed:   time.Since(start),
		Command:   commandWithArgs,
	}
}
