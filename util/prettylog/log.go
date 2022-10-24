package prettylog

import (
	"fmt"
	"time"
)

type logBody struct {
	kind   string
	time   string
	detail string
}

func Errorf(template string, args ...interface{}) {
	t := time.Now()
	s := fmt.Sprintf(template, args...)
	n := logBody{
		kind:   "ERROR",
		time:   t.Format("2006-01-04 15:04:15"),
		detail: s,
	}
	fmt.Println(n)
}

func Infof(template string, args ...interface{}) {
	t := time.Now()
	s := fmt.Sprintf(template, args...)
	n := logBody{
		kind:   "INFO",
		time:   t.Format("2006-01-04 15:04:15"),
		detail: s,
	}
	fmt.Println(n)
}
