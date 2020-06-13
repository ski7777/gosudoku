package grid

type ExtendedCell struct {
	Cell         *Cell
	AllowedVals  []int
	NAllowedVals int
	Used         bool
}

func (ec *ExtendedCell) Update() {
	ec.AllowedVals = ec.Cell.GetAllowedVals()
	ec.NAllowedVals = len(ec.AllowedVals)
	ec.Used = ec.Cell.GetValue() != nil
}

func NewExtendedCell(c *Cell) *ExtendedCell {
	ec := new(ExtendedCell)
	ec.Cell = c
	return ec
}
