package output

import "github.com/ski7777/gosudoku/package/grid"

type Outputter interface {
	Output(*grid.Grid)
}
