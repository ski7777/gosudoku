package output

import (
	"log"
	"sync"

	"github.com/ski7777/gosudoku/package/grid"
)

type Output struct {
	outputters []Outputter
	lock       sync.Mutex
	count      int
	limit      int
	limitcall  func()
}

func (o *Output) RegisterOutputter(op Outputter) {
	o.outputters = append(o.outputters, op)
}

func (o *Output) Output(g *grid.Grid) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.count == o.limit+1 {
		defer func() { go o.limitcall() }()
		return
	}
	log.Println("Solution", o.count)
	for _, op := range o.outputters {
		op.Output(g)
	}
	o.count++
}

func NewOutput(limit int, limitcall func()) *Output {
	o := new(Output)
	o.limit = limit
	o.limitcall = limitcall
	o.count = 1
	return o
}
