package processor

import (
	"HashChallenge/external"
	interfaces "HashChallenge/interfaces"
	"HashChallenge/logger"
	"HashChallenge/util"
	"fmt"
	"net/url"
	"sync"
)

type RequestProcessor struct {
	inputData interfaces.IncomingUrlsRequest
}

func (r *RequestProcessor) Init(parallelCount int, urls []string) error {
	// Prepare input struct
	return r.preprocessInputStruct(parallelCount, &urls)

}
func (r *RequestProcessor) Process() {
	if r == nil || len(r.inputData.Urls) == 0 {
		return
	}

	// Now, Make http calls, create MD5 hash
	response, err := r.makeAPICalls()
	if err == nil {
		printResponse(&response)
	}
}
func printResponse(response *[]interfaces.ResponseStruct) {
	for _, res := range *response {
		fmt.Print(res.InputUrl)
		fmt.Print(" ")
		fmt.Println(res.HashedOutput)
	}
}

func (r *RequestProcessor) preprocessInputStruct(parallelCount int, inputUrls *[]string) error {
	input := interfaces.IncomingUrlsRequest{}
	input.ParallelGoRoutineLimit = parallelCount
	var processedUrls []*url.URL
	for _, url := range *inputUrls {
		// Assumption: In case of URL parsing fails, we will ignore it. We can break the execution too.!
		processedUrl, err := util.ValidateAndParseURL(url)
		if err == nil {
			processedUrls = append(processedUrls, processedUrl)
		} else {
			logger.PushErrLogs("Invalid URL: "+url, logger.Level.Error, err)
		}
	}
	input.Urls = processedUrls
	r.inputData = input
	return nil
}

// makeAPICalls:= Process all the input Urls and receive hashed output along with it.
func (r *RequestProcessor) makeAPICalls() ([]interfaces.ResponseStruct, error) {
	var response []interfaces.ResponseStruct
	var wg sync.WaitGroup

	// Need to limit the number of maximum spawned go-routines.
	c := make(chan interface{}, r.inputData.ParallelGoRoutineLimit)
	for _, eachUrl := range r.inputData.Urls {
		wg.Add(1)
		c <- struct{}{}
		go func(url *url.URL) {
			defer wg.Done()
			res, err := external.GetUrlResponse(url)
			if err == nil {
				response = append(response, res)
			} else {
				// Suppressing the error.
				logger.PushErrLogs("Error occurred while making a get call: "+url.String(), logger.Level.Error, err)
			}
			<-c
		}(eachUrl)
	}
	// Wait till all the responses are received.
	wg.Wait()
	defer close(c)
	// TODO: return proper error if any.
	return response, nil
}
