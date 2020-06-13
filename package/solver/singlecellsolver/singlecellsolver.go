package singlecellsolver

import (
	"github.com/ski7777/gosudoku/package/grid"
)

type SingleCellSolver struct {
}

func (s *SingleCellSolver) Solve(g *grid.Grid) bool {
	changed := false
	cells := g.Cells()
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			c := cells[x][y]
			if c.GetValue() == nil {
				a := c.GetAllowedVals()
				if len(a) == 1 {
					c.SetValue(a[0])
					changed = true
				}
			}
		}
	}
	return changed
}

func Init() *SingleCellSolver {
	s := new(SingleCellSolver)
	return s
}
