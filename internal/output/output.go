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
	o.count++
	if o.count == o.limit {
		defer func() { go o.limitcall() }()
	}
	log.Println("Solution", o.count)
	for _, op := range o.outputters {
		op.Output(g)
	}
}

func NewOutput(limit int, limitcall func()) *Output {
	o := new(Output)
	o.limit = limit
	o.limitcall = limitcall
	o.count = 0
	if o.limit == 0 {
		go o.limitcall()
	}
	return o
}
