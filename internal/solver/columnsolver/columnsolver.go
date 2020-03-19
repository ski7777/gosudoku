package columnsolver

import (
	"github.com/ski7777/gosudoku/internal/grid"
)

type ColumnSolver struct {
}

func (s *ColumnSolver) Solve(g *grid.Grid) bool {
	changed := false
	cols := g.Cols()
	for y := 0; y < 9; y++ {
		c := cols[y]
		ccs := c.Cells()
		a := c.GetAllowedVals()
		for i := 0; i < len(a); i++ {
			v := a[i]
			ac := make([]*grid.Cell, 0)
			for x := 0; x < 9; x++ {
				cc := ccs[x]
				if cc.GetValue() == nil {
					cca := cc.GetAllowedVals()
					for j := 0; j < len(cca); j++ {
						if cca[j] == v {
							ac = append(ac, cc)
						}
					}
				}
			}
			if len(ac) == 1 {
				ac[0].SetValue(v)
				changed = true
			}
		}
	}
	return changed
}

func Init() *ColumnSolver {
	s := new(ColumnSolver)
	return s
}
