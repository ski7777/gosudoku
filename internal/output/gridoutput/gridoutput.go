package gridoutput

import (
	"github.com/ski7777/gosudoku/package/grid"
)

type GridOutput struct{}

func (GridOutput) Output(g *grid.Grid) {
	g.Print()
}
