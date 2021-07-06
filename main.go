package main

import (
	"HashChallenge/constants"
	"HashChallenge/logger"
	"HashChallenge/processor"
	"errors"
	"flag"
	"strconv"
)

func main() {
	parallelLimit := flag.Int("parallel", constants.DEFAULT_MAX_PARALLEL_LIMIT, "Maximum number of concurrent requests")
	flag.Parse()
	logger.PushLogs("Maximum number of concurrent requests Limit - Parallel:"+strconv.Itoa(*parallelLimit), logger.Level.Info)
	arguements := flag.Args()
	if len(arguements) == 0 {
		logger.PushErrLogs("Invalid Input - ", logger.Level.Error, errors.New("Invalid arguments"))
		return
	}
	requestProcessor := processor.RequestProcessor{}
	requestProcessor.Init(*parallelLimit, flag.Args())
	requestProcessor.Process()
}
