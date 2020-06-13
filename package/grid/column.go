package grid

type Column struct {
	cells [9]*Cell
}

func (c *Column) GetVals() []int {
	var vals []int
	for i := 0; i < 9; i++ {
		if v := c.cells[i].GetValue(); v != nil {
			vals = append(vals, *v)
		}
	}
	return vals
}

func (c *Column) Cells() [9]*Cell {
	return c.cells
}

func (c *Column) GetAllowedVals() []int {
	allowed := make([]int, 0)
	vals := c.GetVals()
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

func newCol(cells [9]*Cell) *Column {
	c := new(Column)
	c.cells = cells
	return c
}
