package solvermanager

import (
	"github.com/ski7777/gosudoku/package/grid"
)

type Solver interface {
	Solve(*grid.Grid) bool
}
