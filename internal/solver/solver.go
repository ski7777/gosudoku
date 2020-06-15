package solver

import (
	"sort"
	"sync"

	"github.com/ski7777/gosudoku/internal/workermanager"
	"github.com/ski7777/gosudoku/package/grid"
)

type Solver struct {
	wm              *workermanager.WorkerManager
	resultCall      func(*grid.Grid)
	grids, oldgrids [82][]*grid.ExtendedGrid
	gridlock        sync.Mutex
	endcall         func()
	update          *sync.Cond
	stats           *[3]int
	newstats        *sync.Cond
}

func (sm *Solver) workerManagerCallBack(data workermanager.Result) {
	sm.gridlock.Lock()
	defer func() {
		sm.update.L.Lock()
		sm.update.Broadcast()
		sm.update.L.Unlock()
	}()
	defer sm.gridlock.Unlock()
inloop:
	for _, eg := range data {
		for _, og := range sm.oldgrids[eg.GetFreeCount()] {
			if og.GetAllowedCount() == eg.GetAllowedCount() {
				if og.GetGrid().Equals(eg.GetGrid()) {
					continue inloop
				}
			}
		}
		//gird is new. yay
		sm.grids = appendGridMatrix(sm.grids, eg)
		sm.oldgrids = appendGridMatrix(sm.oldgrids, eg)
		if eg.GetFreeCount() == 0 {
			sm.resultCall(eg.GetGrid())
		}
	}
}

func (sm *Solver) Solve() {
	sm.gridlock.Lock()
	var wmn, wmi, wmw int
	sm.wm.ForcePushStats()
	wmn = sm.stats[0]
solve:
	for {
		n := 0
		for x := 1; x < 82; x++ {
			n += len(sm.grids[x])
		}
	job:
		for i := 0; i < wmn; i++ {
			for x := 1; x < 82; x++ {
				if len(sm.grids[x]) > 0 {
					sm.wm.Work(sm.grids[x][0])
					sm.grids[x] = sm.grids[x][1:]
					continue job
				}
			}
			break job
		}
		sm.gridlock.Unlock()
	finish:
		for {
			sm.update.L.Lock()
			sm.update.Wait()
			sm.update.L.Unlock()
			sm.newstats.L.Lock()
			sm.newstats.Wait()
			sm.newstats.L.Unlock()
			wmi, wmw = sm.stats[1], sm.stats[2]
			if wmi == 0 && wmw == 0 {
				break finish
			}
		}
		sm.gridlock.Lock()
		for _, gl := range sm.grids[1:] {
			if len(gl) > 0 {
				continue solve
			}
		}
		sm.gridlock.Unlock()
		sm.endcall()
		return
	}
}

func NewSolver(g *grid.Grid, wm *workermanager.WorkerManager, resultCall func(*grid.Grid), endcall func()) *Solver {
	sm := new(Solver)
	sm.wm = wm
	sm.wm.SetCallback(sm.workerManagerCallBack)
	sm.resultCall = resultCall
	sm.endcall = endcall
	sm.update = sync.NewCond(&sync.Mutex{})
	eg := grid.NewExtendedGrid(g)
	eg.UpdateAll()
	sm.workerManagerCallBack(workermanager.Result{eg})
	sm.stats, sm.newstats = sm.wm.GetStats()
	return sm
}

func appendGridMatrix(data [82][]*grid.ExtendedGrid, n *grid.ExtendedGrid) [82][]*grid.ExtendedGrid {
	worklist := append(data[n.GetFreeCount()], n)
	sort.Slice(worklist, func(i, j int) bool {
		return worklist[i].GetAllowedCount() < worklist[j].GetAllowedCount()
	})
	data[n.GetFreeCount()] = worklist
	return data
}
