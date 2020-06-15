package output

import "github.com/ski7777/gosudoku/package/grid"

type GridOutputter interface {
	Output(*grid.Grid)
}
