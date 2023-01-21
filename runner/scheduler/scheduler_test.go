package scheduler

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseScheduleFromCommand(t *testing.T) {

	type test struct {
		cronCommand string
		schedule    string
		command     string
		isError     bool
		nextRunTime string
	}

	timeFrom, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:00:12")

	testTable := []test{
		{cronCommand: "* * * * * php -v", schedule: "* * * * *", command: "php -v", isError: false, nextRunTime: "2023-01-01 12:01:00"},
		{cronCommand: "*/1 * * * * php -v", schedule: "*/1 * * * *", command: "php -v", isError: false, nextRunTime: "2023-01-01 12:01:00"},
		{cronCommand: "1 13 * * * /project/bin/symfony command:action arg 123", schedule: "1 13 * * *", command: "/project/bin/symfony command:action arg 123", isError: false, nextRunTime: "2023-01-01 13:01:00"},
		{cronCommand: "0 2 * * * /test-exec", schedule: "0 2 * * *", command: "/test-exec", isError: false, nextRunTime: "2023-01-02 02:00:00"},
		{cronCommand: "0 2 * * php -v", schedule: "0 2 * * *", command: "php -v", isError: true, nextRunTime: "2023-01-02 02:00:00"},
	}

	for _, testCase := range testTable {
		commandSchedule, err := ParseScheduleFromCommand(testCase.cronCommand)

		if testCase.isError {
			assert.NotNil(t, err)
			continue
		}

		assert.Equal(t, testCase.schedule, commandSchedule.ScheduleString)
		assert.Equal(t, testCase.command, commandSchedule.CommandString)
		assert.Equal(t, testCase.nextRunTime, commandSchedule.CronExpression.Next(timeFrom).Format("2006-01-02 15:04:05"))
	}
}
