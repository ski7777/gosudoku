package algorithmic

import (
	"github.com/ski7777/gosudoku/internal/grid"
)

type Solver interface {
	Solve(*grid.Grid) bool
}
