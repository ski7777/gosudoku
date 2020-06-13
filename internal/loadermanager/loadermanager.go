package loadermanager

import (
	"errors"
	"strings"

	"github.com/ski7777/gosudoku/internal/loader"
	"github.com/ski7777/gosudoku/package/grid"
)

type LoadermManager struct {
	loaders map[string]loader.Loader
	args    map[string]string
}

func (lm *LoadermManager) AddLoader(l loader.Loader) {
	lm.loaders[l.GetShortName()] = l
}

func (lm *LoadermManager) GetLoaders() map[string]loader.Loader {
	return lm.loaders
}

func (lm *LoadermManager) Load(name string, g *grid.Grid) error {
	if l, ok := lm.loaders[name]; ok {
		return l.Load(g, lm.args)
	} else {
		return errors.New("Loader not found")
	}
}

func (lm *LoadermManager) ParseArgs(raw []string) {
	for i := 0; i < len(raw); i++ {
		ra := strings.Split(raw[i], "=")
		lm.args[ra[0]] = ra[1]
	}
}

func NewLoadermManager() *LoadermManager {
	lm := new(LoadermManager)
	lm.loaders = make(map[string]loader.Loader)
	lm.args = make(map[string]string)
	return lm
}
