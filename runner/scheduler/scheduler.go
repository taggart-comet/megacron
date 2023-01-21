package scheduler

import (
	"errors"
	"github.com/gorhill/cronexpr"
	"megacron/system/functions"
	"strings"
)

type CommandSchedule struct {
	OriginalString string
	ScheduleString string
	CommandString  string
	CronExpression *cronexpr.Expression
}

func ParseScheduleFromCommand(cronCommand string) (CommandSchedule, error) {

	commandSchedule := CommandSchedule{OriginalString: cronCommand}

	parts := functions.SplitString(cronCommand, " ")
	if len(parts) < 6 {
		return commandSchedule, errors.New("scheduler: invalid cron command format")
	}
	commandSchedule.ScheduleString = strings.Join(parts[0:5], " ")
	commandSchedule.CommandString = strings.Join(parts[5:], " ")

	cronExpression, cronParseError := cronexpr.Parse(commandSchedule.ScheduleString)

	if cronParseError != nil {
		return commandSchedule, cronParseError
	}

	commandSchedule.CronExpression = cronExpression

	return commandSchedule, nil
}
