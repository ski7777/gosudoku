package grid

type Block struct {
	cells [3][3]*Cell
}

func (b *Block) GetVals() []int {
	var vals []int
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if v := b.cells[x][y].GetValue(); v != nil {
				vals = append(vals, *v)
			}
		}
	}
	return vals
}

func (b *Block) GetAllowedVals() []int {
	allowed := make([]int, 0)
	vals := b.GetVals()
	for i := 1; i <= 9; i++ {
		found := false
		for j := 0; j < len(vals); j++ {
			if vals[j] == i {
				found = true
			}
		}
		if !found {
			allowed = append(allowed, i)
		}
	}
	return allowed
}

func (b *Block) Cells() [3][3]*Cell {
	return b.cells
}

func newBlock(cells [3][3]*Cell) *Block {
	b := new(Block)
	b.cells = cells
	return b
}
