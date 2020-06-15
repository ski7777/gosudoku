package workermanager

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ski7777/gosudoku/package/grid"
)

type WorkerManager struct {
	workers     []Worker
	workerslock *sync.Mutex
	newworkers  *sync.Cond
	inWork      map[string]*job
	inWorkLock  sync.Mutex
	waiting     int
	waitinglock sync.Mutex
	callback    CallBack
	stats       *[3]int
	newstats    *sync.Cond
}

func (wm *WorkerManager) AddWorker(wl ...Worker) {
	wm.workerslock.Lock()
	defer wm.workerslock.Unlock()
	for _, nw := range wl {
		wm.workers = append(wm.workers, nw)
		wm.newworkers.Signal()
	}
}

func (wm *WorkerManager) Work(eg *grid.ExtendedGrid) error {
	j := &job{eg, time.Now()}
	wm.waitinglock.Lock()
	wm.waiting++
	wm.waitinglock.Unlock()
	var id string
	if uuid, err := uuid.NewRandom(); err != nil {
		return err
	} else {
		id = uuid.String()
	}
	wm.inWorkLock.Lock()
	wm.inWork[id] = j
	wm.inWorkLock.Unlock()
	wm.waitinglock.Lock()
	wm.waiting--
	wm.waitinglock.Unlock()
	var w Worker
wait:
	for {
		wm.workerslock.Lock()
		if len(wm.workers) > 0 {
			w = wm.workers[0]
			wm.workers = wm.workers[1:]
			break wait
		}
		wm.newworkers.Wait()
		wm.workerslock.Unlock()
	}
	wm.workerslock.Unlock()
	go func() {
		r := w(j.eg)
		wm.inWorkLock.Lock()
		defer wm.inWorkLock.Unlock()
		if _, ok := wm.inWork[id]; ok {
			delete(wm.inWork, id)
			go wm.callback(r)
			wm.AddWorker(w)
			wm.workerslock.Lock()
			defer wm.workerslock.Unlock()
			wm.waitinglock.Lock()
			defer wm.waitinglock.Unlock()
			wm.pushStats()
		}
	}()
	return nil
}

func (wm *WorkerManager) SetCallback(cb CallBack) {
	wm.callback = cb
}

//no locking! Needs workerslock, inWorkLock, waitinglock
//non-blocking!
func (wm *WorkerManager) pushStats() {
	wm.newstats.L.Lock()
	defer wm.newstats.L.Unlock()
	wm.stats[0], wm.stats[1], wm.stats[2] = len(wm.workers), len(wm.inWork), wm.waiting
	wm.newstats.Broadcast()
}

func (wm *WorkerManager) ForcePushStats() {
	wm.workerslock.Lock()
	wm.inWorkLock.Lock()
	wm.waitinglock.Lock()
	defer wm.workerslock.Unlock()
	defer wm.inWorkLock.Unlock()
	defer wm.waitinglock.Unlock()
	wm.pushStats()
}

func (wm *WorkerManager) GetStats() (*[3]int, *sync.Cond) {
	return wm.stats, wm.newstats
}

func NewWorkerManager() *WorkerManager {
	wm := new(WorkerManager)
	wm.workerslock = &sync.Mutex{}
	wm.newworkers = sync.NewCond(wm.workerslock)
	wm.inWork = make(map[string]*job)
	wm.stats = new([3]int)
	wm.newstats = sync.NewCond(&sync.Mutex{})
	go func() {
		for {
			<-time.Tick(time.Second)
			wm.ForcePushStats()
		}
	}()
	return wm
}
