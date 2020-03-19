package solvermanager

import (
	"log"

	"github.com/ski7777/gosudoku/internal/grid"
	stringloader "github.com/ski7777/gosudoku/internal/loader/string"
)

type SolverManager struct {
	grid      *grid.Grid
	solver    []Solver
}

func (sm *SolverManager) SolveAlgorithmic(g *grid.Grid) bool {
	if g == nil {
		g = sm.grid
	}
	var changed bool
	for {
		for si := 0; si < len(sm.solver); si++ {
			s := sm.solver[si]
			if s.Solve(g) {
				changed = true
			}
			if sm.grid.CheckSolved() {
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

func (sm *SolverManager) clone() *grid.Grid {
	g := grid.NewGrid()
	stringloader.Init().Load(g, map[string]string{"data": sm.grid.ToString()})
	return g
}

func NewSolverManager(grid *grid.Grid) *SolverManager {
	sm := new(SolverManager)
	sm.grid = grid
	sm.solver = append(sm.solver,
	)
	return sm
}
