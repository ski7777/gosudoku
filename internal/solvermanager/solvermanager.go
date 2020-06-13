package solvermanager

import (
	"log"
	"os"
	"sort"
	"sync"

	"github.com/ski7777/gosudoku/package/algorithmicsolver"
	"github.com/ski7777/gosudoku/package/grid"
)

type extGrid struct {
	grid *grid.Grid
	free int
	n    int
	sec  []*grid.ExtendedCell
}

func (eg *extGrid) updateN() {
	n := 0
	for _, m := range eg.sec {
		n += m.NAllowedVals
	}
}

func (eg *extGrid) updateSec() {
	eg.sec = eg.grid.GetSortedExtendedCells(false)
}

func (eg *extGrid) updateAll() {
	eg.free = eg.grid.GetNumFree()
	eg.updateSec()
	eg.updateN()
}

type SolverManager struct {
	grid                    *grid.Grid
	as                      *algorithmicsolver.AlgorithmicSolver
	str                     bool
	workerproc              int
	workerproclock          sync.Mutex
	steps                   int
	stepslock               sync.Mutex
	sol                     int
	sollock                 sync.Mutex
	maxprocs, maxsolutions  int
	grids, oldgrids         [82][]*extGrid
	gridslock, olfgridslock sync.Mutex
	workers                 chan func(*extGrid)
	workerslock             sync.Mutex
	printlock               sync.Mutex
	println                 bool
	printwait               sync.WaitGroup
	end                     chan interface{}
	init                    bool
	waiting                 int
	waitinglock             sync.Mutex
}

func (sm *SolverManager) push(g *extGrid) {
	sm.printwait.Add(1)
	go func() {
		sm.stepslock.Lock()
		sm.printlock.Lock()
		sm.gridslock.Lock()
		sm.workerproclock.Lock()
		sm.workerslock.Lock()
		sm.waitinglock.Lock()
		defer sm.stepslock.Unlock()
		defer sm.printlock.Unlock()
		defer sm.gridslock.Unlock()
		defer sm.workerproclock.Unlock()
		defer sm.workerslock.Unlock()
		defer sm.waitinglock.Unlock()
		sm.steps++
		pending := sm.workerproc + (sm.maxprocs/2 - len(sm.workers)) + sm.waiting
		for _, gl := range sm.grids[1:] {
			pending += len(gl)
		}
		print("\r                         \rStep ", sm.steps, ", ", pending, " pending")
		sm.println = true
		if pending == 0 && sm.init {
			sm.end <- struct{}{}
		}
	}()
	sm.olfgridslock.Lock()
	defer sm.olfgridslock.Unlock()
	for _, og := range sm.oldgrids[g.free] {
		if og.n == g.n {
			if og.grid.Equals(g.grid) {
				return
			}
		}
	}
	worklist := append(sm.oldgrids[g.free], g)
	sort.Slice(worklist, func(i, j int) bool {
		return worklist[i].n < worklist[j].n
	})
	sm.oldgrids[g.free] = worklist
	go func() {
		if g.free == 0 {
			sm.print(g.grid)
		}
		sm.printwait.Done()
	}()
	sm.gridslock.Lock()
	defer sm.gridslock.Unlock()
	worklist = append(sm.grids[g.free], g)
	sort.Slice(worklist, func(i, j int) bool {
		return worklist[i].n < worklist[j].n
	})
	sm.grids[g.free] = worklist
}

func (sm *SolverManager) Solve() {
	sm.as.Solve(sm.grid)
	eg := &extGrid{grid: sm.grid}
	eg.updateAll()
	sm.push(eg)
	sm.init = true
	sm.workerslock.Lock()
	for i := 0; i < (sm.maxprocs / 2); i++ {
		sm.workers <- sm.worker
	}
	sm.workerslock.Unlock()
	go func() {
		for {
			sm.gridslock.Lock()
			var job *extGrid
			for i := 1; i < 82 && job == nil; i++ {
				l := sm.grids[i]
				if len(l) > 0 {
					job = l[0]
					sm.grids[i] = sm.grids[i][1:]
				}
			}
			sm.gridslock.Unlock()
			if job == nil {
				continue
			}
			sm.workerslock.Lock()
			select {
			case w := <-sm.workers:
				{
					go w(job)
				}
			default:
				{
				}
			}
			sm.workerslock.Unlock()
		}
	}()
	<-sm.end
	sm.printwait.Wait()
	if sm.println {
		println()
	}
}

func (sm *SolverManager) worker(eg *extGrid) {
	for _, ec := range eg.sec {
		c := ec.Cell
		sm.waitinglock.Lock()
		sm.waiting += len(ec.AllowedVals)
		sm.waitinglock.Unlock()
		for _, s := range ec.AllowedVals {
			c.SetValue(s)
			ng := eg.grid.Clone()
			go func() {
			waitloop:
				for {
					sm.workerproclock.Lock()
					if sm.workerproc < (sm.maxprocs / 2) {
						sm.workerproc++
						sm.workerproclock.Unlock()
						sm.waitinglock.Lock()
						sm.waiting--
						sm.waitinglock.Unlock()
						solved := sm.as.Solve(ng)
						neg := &extGrid{grid: ng}
						if solved {
							neg.free = 0
							neg.n = 0
						} else {
							neg.updateAll()
						}
						go func() {
							sm.push(neg)
							sm.workerproclock.Lock()
							sm.workerproc--
							sm.workerproclock.Unlock()
						}()
						break waitloop
					} else {
						sm.workerproclock.Unlock()
					}
				}
			}()
		}
		c.UnsetValue()
	}
	sm.workerslock.Lock()
	defer sm.workerslock.Unlock()
	sm.workers <- sm.worker
}

func (sm *SolverManager) print(g *grid.Grid) {
	sm.printlock.Lock()
	defer sm.printlock.Unlock()
	if sm.println {
		println()
		sm.println = false
	}
	sm.sollock.Lock()
	defer sm.sollock.Unlock()
	log.Println("Solution", sm.sol)
	g.Print()
	if sm.str {
		log.Println("String:", g.ToString())
	}
	sm.sol++
	if sm.sol == sm.maxsolutions {
		os.Exit(0)
	}
}

func NewSolverManager(grid *grid.Grid, str bool, maxprocs, maxsolutions int) *SolverManager {
	sm := new(SolverManager)
	sm.grid = grid
	sm.str = str
	sm.sol = 1
	sm.as = algorithmicsolver.NewAlgorithmicSolver()
	sm.maxprocs = maxprocs
	sm.maxsolutions = maxsolutions
	sm.workers = make(chan func(*extGrid), sm.maxprocs/2)
	sm.end = make(chan interface{})
	return sm
}
