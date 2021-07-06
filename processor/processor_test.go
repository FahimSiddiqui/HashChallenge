package processor

import (
	"testing"
)

func TestRequestInit(t *testing.T) {
	maxParallelLimit := 2
	requestProcessor := RequestProcessor{}
	requestProcessor.Init(maxParallelLimit, []string{"google.com", "test.ai", "http://facebook.com"})
	if requestProcessor.inputData.ParallelGoRoutineLimit != maxParallelLimit && len(requestProcessor.inputData.Urls) != 3 {
		t.Error()
	}
}

func TestAPICalls(t *testing.T) {
	maxParallelLimit := 2
	requestProcessor := RequestProcessor{}
	requestProcessor.Init(maxParallelLimit, []string{"google.com", "test.ai", "http://facebook.com"})

	_, err := requestProcessor.makeAPICalls()
	if err != nil {
		t.Fail()
	}
}
