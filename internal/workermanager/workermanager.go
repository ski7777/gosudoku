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
		}
	}()
	return nil
}

func (wm *WorkerManager) SetCallback(cb CallBack) {
	wm.callback = cb
}

func NewWorkerManager() *WorkerManager {
	wm := new(WorkerManager)
	wm.workerslock = &sync.Mutex{}
	wm.newworkers = sync.NewCond(wm.workerslock)
	return wm
}
