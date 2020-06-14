package workermanager

import (
	"time"

	"github.com/ski7777/gosudoku/package/grid"
)

type job struct {
	eg      *grid.ExtendedGrid
	started time.Time
}
