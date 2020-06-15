package solver

import (
	"github.com/ski7777/gosudoku/internal/workermanager"
	"github.com/ski7777/gosudoku/package/grid"
	"github.com/ski7777/gosudoku/package/solvermanager"
)

func GetSolverWorker(s solvermanager.Solver) workermanager.Worker {
	return func(eg *grid.ExtendedGrid) (r workermanager.Result) {
		if solved := s.Solve(eg.GetGrid()); solved {
			eg.SetNull()
			r = append(r, eg)
			return
		}
		eg.UpdateAll()
		for _, ec := range eg.GetSEC() {
			c := ec.GetCell()
			for _, v := range ec.GetAllowedValues() {
				c.SetValue(v)
				neg := eg.Clone()
				neg.UpdateAll()
				r = append(r, neg)
			}
			c.UnsetValue()
		}
		return
	}
}
