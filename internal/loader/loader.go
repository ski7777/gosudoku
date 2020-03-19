package loader

import (
	"github.com/ski7777/gosudoku/internal/grid"
)

type Loader interface {
	Load(*grid.Grid, map[string]string) error
	GetShortName() string
	GetLongName() string
	GetDescription() string
}
