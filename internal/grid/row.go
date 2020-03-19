package grid

type Row struct {
	cells [9]*Cell
}

func (r *Row) GetVals() []int {
	var vals []int
	for i := 0; i < 9; i++ {
		if v := r.cells[i].GetValue(); v != nil {
			vals = append(vals, *v)
		}
	}
	return vals
}

func (r *Row) Cells() [9]*Cell {
	return r.cells
}

func (r *Row) GetAllowedVals() []int {
	allowed := make([]int, 0)
	vals := r.GetVals()
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

func newRow(cells [9]*Cell) *Row {
	r := new(Row)
	r.cells = cells
	return r
}
