package prettylog

import (
	"fmt"
	"time"
)

func Errorf(template string, args ...interface{}) {
	t := time.Now()
	s := fmt.Sprintf(template, args)
	fmt.Println("[  " + t.Format("2006-01-04 15:04:15") + "    ERROR   " + "   " + s + "  ]")
}

func Infof(template string, args ...interface{}) {
	t := time.Now()
	s := fmt.Sprintf(template, args)
	fmt.Println("[  " + t.Format("2006-01-04 15:04:15") + "    INFO   " + "   " + s + "  ]")
}
