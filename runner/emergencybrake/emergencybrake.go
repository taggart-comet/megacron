package emergencybrake

import (
	"megacron/system/functions"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type EmergencyBrake struct {
	shouldStop   bool
	waitingGroup *sync.WaitGroup
}

type EmergencyBrakeInterface interface {
	ShouldStop() bool
	Start()
	Stop()
	Wait()
}

func NewEmergencyBrake() EmergencyBrakeInterface {
	eb := EmergencyBrake{}
	eb.shouldStop = false
	eb.waitingGroup = &sync.WaitGroup{}

	go func() {

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGHUP)

		// Block until we receive our signal.
		<-c

		eb.shouldStop = true
	}()

	return &eb
}

func (eb *EmergencyBrake) ShouldStop() bool {
	return eb.shouldStop
}

func (eb *EmergencyBrake) Start() {
	eb.waitingGroup.Add(1)
}

func (eb *EmergencyBrake) Stop() {
	eb.waitingGroup.Done()
}

func (eb *EmergencyBrake) Wait() {

	functions.Log("Waiting for all crons to finish on their own...")
	eb.waitingGroup.Wait()
}
