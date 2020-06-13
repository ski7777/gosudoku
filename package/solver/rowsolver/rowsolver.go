package rowsolver

import (
	"github.com/ski7777/gosudoku/package/grid"
)

type RowSolver struct {
}

func (s *RowSolver) Solve(g *grid.Grid) bool {
	changed := false
	rows := g.Rows()
	for y := 0; y < 9; y++ {
		r := rows[y]
		rcs := r.Cells()
		a := r.GetAllowedVals()
		for i := 0; i < len(a); i++ {
			v := a[i]
			ac := make([]*grid.Cell, 0)
			for x := 0; x < 9; x++ {
				rc := rcs[x]
				if rc.GetValue() == nil {
					rca := rc.GetAllowedVals()
					for j := 0; j < len(rca); j++ {
						if rca[j] == v {
							ac = append(ac, rc)
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

func Init() *RowSolver {
	s := new(RowSolver)
	return s
}
