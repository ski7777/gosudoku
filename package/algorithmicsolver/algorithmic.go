package algorithmicsolver

import (
	"github.com/ski7777/gosudoku/package/grid"
	"github.com/ski7777/gosudoku/package/solver/blocksolver"
	"github.com/ski7777/gosudoku/package/solver/columnsolver"
	"github.com/ski7777/gosudoku/package/solver/rowsolver"
	"github.com/ski7777/gosudoku/package/solver/singlecellsolver"
	"github.com/ski7777/gosudoku/package/solvermanager"
)

type AlgorithmicSolver struct {
	solver []solvermanager.Solver
}

func (as *AlgorithmicSolver) Solve(g *grid.Grid) bool {
	if g == nil {
		panic("g must not be nil")
	}
	var changed bool
	for {
		for si := 0; si < len(as.solver); si++ {
			s := as.solver[si]
			if s.Solve(g) {
				changed = true
			}
			if g.CheckSolved() {
				return true
			}
		}
		if changed {
			changed = false
		} else {
			return false
		}
	}
}

func NewAlgorithmicSolver() *AlgorithmicSolver {
	as := new(AlgorithmicSolver)
	as.solver = append(as.solver,
		singlecellsolver.Init(),
		blocksolver.Init(),
		rowsolver.Init(),
		columnsolver.Init(),
	)
	return as
}
