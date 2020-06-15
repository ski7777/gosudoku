package output

import (
	"log"
	"sync"

	"github.com/ski7777/gosudoku/package/grid"
)

type Output struct {
	gridoutputters []GridOutputter
	logoutputter   LogOutputter
	lock           sync.Mutex
	count          int
	limit          int
	limitcall      func()
	lastlog        bool
}

func (o *Output) RegisterGridOutputter(op GridOutputter) {
	o.gridoutputters = append(o.gridoutputters, op)
}

func (o *Output) SetLogOutputter(op LogOutputter) {
	o.logoutputter = op
}

func (o *Output) OutputGrid(g *grid.Grid) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.lastlog && o.logoutputter != nil {
		o.logoutputter.CleanLine()
	}
	o.count++
	if o.count == o.limit {
		defer func() { go o.limitcall() }()
	}
	log.Println("Solution", o.count)
	for _, op := range o.gridoutputters {
		op.Output(g)
	}
}

func (o *Output) OutputLog(l ...interface{}) {
	if o.logoutputter != nil {
		o.lock.Lock()
		defer o.lock.Unlock()
		o.logoutputter.Output(l)
		o.lastlog = true
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
