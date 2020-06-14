package stringoutput

import (
	"github.com/ski7777/gosudoku/package/grid"
)

type StringOutput struct{}

func (StringOutput) Output(g *grid.Grid) {
	println("String:", g.ToString())
}
