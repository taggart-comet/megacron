package runner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOkRun(t *testing.T) {
	feedBackChannel := make(chan RunResult, 1)
	defer close(feedBackChannel)

	runner := Runner{}
	runner.Run("echo ok", feedBackChannel)

	runResult := <-feedBackChannel

	assert.Equal(t, "ok", runResult.Output, "output is not correct")
	assert.Equal(t, true, runResult.IsSuccess, "status is not correct")
}

func TestErrorRun(t *testing.T) {
	feedBackChannel := make(chan RunResult, 1)
	defer close(feedBackChannel)

	runner := Runner{}
	runner.Run("exit 1", feedBackChannel)

	runResult := <-feedBackChannel

	assert.Equal(t, "", runResult.Output, "output is not correct")
	assert.Equal(t, false, runResult.IsSuccess, "status is not correct")
}
