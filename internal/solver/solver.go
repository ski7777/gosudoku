package solver

import (
	"github.com/ski7777/gosudoku/internal/workermanager"
	"github.com/ski7777/gosudoku/package/grid"
)

type Solver struct {
	wm              *workermanager.WorkerManager
	resultCall      func(*grid.Grid)
	grids, oldgrids [82][]*grid.ExtendedGrid
	endcall         func()
}

func (sm *Solver) workerManagerCallBack(data workermanager.Result) {}

func (sm *Solver) Solve() {}

func NewSolver(g *grid.Grid, wm *workermanager.WorkerManager, resultCall func(*grid.Grid), endcall func()) *Solver {
	sm := new(Solver)
	sm.wm = wm
	sm.wm.SetCallback(sm.workerManagerCallBack)
	sm.resultCall = resultCall
	sm.endcall = endcall
	return sm
}