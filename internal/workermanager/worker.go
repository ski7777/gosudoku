package workermanager

import "github.com/ski7777/gosudoku/package/grid"

type Worker func(*grid.ExtendedGrid) Result
