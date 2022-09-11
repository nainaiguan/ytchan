package prettylog

import (
	"fmt"
	"time"
)

func Errorf(template string, args ...interface{}) {
	t := time.Now()
	s := fmt.Sprintf(template, args)
	fmt.Println("[  " + t.String() + "ERROR   " + "   "  + s + "  ]")
}

func Infof(template string, args ...interface{}) {
	t := time.Now()
	s := fmt.Sprintf(template, args)
	fmt.Println("[  " + t.String() + "INFO   " + "   "  + s + "  ]")
}
