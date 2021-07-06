package interfaces

import (
	"net/url"
)

type IncomingUrlsRequest struct {
	ParallelGoRoutineLimit int
	Urls                   []*url.URL
}
