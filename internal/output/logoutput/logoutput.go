package logoutput

import (
	"fmt"
	"strings"
)

type LogOutput struct {
	max int
}

func (o *LogOutput) Output(l []interface{}) {
	o.CleanLine()
	data := []string{}
	for _, e := range l {
		data = append(data, fmt.Sprint(e))
	}
	s := strings.Join(data, " ")
	print(s)
	o.max = len(s)
}

func (o *LogOutput) CleanLine() {
	print("\r")
	for i := 0; i < o.max; i++ {
		print(" ")
	}
	print("\r")
	o.max = 0
}
