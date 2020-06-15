package grid

type ExtendedGrid struct {
	grid *Grid
	free int
	n    int
	sec  []*ExtendedCell
}

func (eg *ExtendedGrid) UpdateN() {
	n := 0
	for _, m := range eg.sec {
		n += m.GetNumAllowedValued()
	}
}

func (eg *ExtendedGrid) UpdateSec() {
	eg.sec = eg.grid.GetSortedExtendedCells(false)
}

func (eg *ExtendedGrid) UpdateAll() {
	eg.free = eg.grid.GetNumFree()
	eg.UpdateSec()
	eg.UpdateN()
}

func (eg *ExtendedGrid) GetGrid() *Grid { return eg.grid }

func (eg *ExtendedGrid) GetFreeCount() int { return eg.free }

func (eg *ExtendedGrid) GetAllowedCount() int { return eg.n }

func (eg *ExtendedGrid) GetSEC() []*ExtendedCell { return eg.sec }

func (eg *ExtendedGrid) SetNull() {
	eg.free = 0
	eg.n = 0
	eg.sec = []*ExtendedCell{}
}

func (eg *ExtendedGrid) Clone() *ExtendedGrid {
	return NewExtendedGrid(eg.GetGrid().Clone())
}

func NewExtendedGrid(g *Grid) *ExtendedGrid {
	eg := new(ExtendedGrid)
	eg.grid = g
	return eg
}
