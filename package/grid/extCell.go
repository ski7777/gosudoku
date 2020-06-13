package grid

type ExtendedCell struct {
	cell         *Cell
	allowedVals  []int
	nAllowedVals int
	used         bool
}

func (ec *ExtendedCell) Update() {
	ec.allowedVals = ec.cell.GetAllowedVals()
	ec.nAllowedVals = len(ec.allowedVals)
	ec.used = ec.cell.GetValue() != nil
}

func (ec *ExtendedCell) getCell() *Cell { return ec.cell }

func (ec *ExtendedCell) getAllowedValues() []int { return ec.allowedVals }

func (ec *ExtendedCell) getNumAllowedValued() int { return ec.nAllowedVals }

func (ec *ExtendedCell) isUsed() bool { return ec.used }

func NewExtendedCell(c *Cell) *ExtendedCell {
	ec := new(ExtendedCell)
	ec.cell = c
	return ec
}
