package ray

import (
	"os"
	"runtime"
)

type origin struct {
	File       string `json:"file"`
	LineNumber int    `json:"line_number"`
	Hostname   string `json:"hostname"`
}

func newOrigin(skip int) *origin {
	o := new(origin)

	_, file, lineNumber, ok := runtime.Caller(skip)
	if ok {
		o.File = file
		o.LineNumber = lineNumber
	}

	hostname, err := os.Hostname()
	if err == nil {
		o.Hostname = hostname
	}

	return o
}
