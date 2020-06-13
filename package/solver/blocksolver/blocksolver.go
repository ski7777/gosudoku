package blocksolver

import (
	"github.com/ski7777/gosudoku/package/grid"
)

type BlockSolver struct {
	grid *grid.Grid
}

func (s *BlockSolver) Solve(g *grid.Grid) bool {
	changed := false
	blocks := g.Blocks()
	for bx := 0; bx < 3; bx++ {
		for by := 0; by < 3; by++ {
			b := blocks[bx][by]
			bcs := b.Cells()
			a := b.GetAllowedVals()
			for i := 0; i < len(a); i++ {
				v := a[i]
				ac := make([]*grid.Cell, 0)
				for x := 0; x < 3; x++ {
					for y := 0; y < 3; y++ {
						bc := bcs[x][y]
						if bc.GetValue() == nil {
							bca := bc.GetAllowedVals()
							for j := 0; j < len(bca); j++ {
								if bca[j] == v {
									ac = append(ac, bc)
								}
							}
						}
					}
				}
				if len(ac) == 1 {
					ac[0].SetValue(v)
					changed = true
				}
			}
		}
	}
	return changed
}

func Init() *BlockSolver {
	s := new(BlockSolver)
	return s
}
