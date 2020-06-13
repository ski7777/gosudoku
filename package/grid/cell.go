package grid

import (
	"errors"
)

type Cell struct {
	row   *Row
	col   *Column
	block *Block
	val   *int
}

func (c *Cell) UnsetValue() {
	c.val = nil
}

func (c *Cell) SetValue(v int) error {
	if v < 1 || v > 9 {
		return errors.New("Value not in allowed range (1-9)")
	}
	c.val = &v
	return nil
}

func (c *Cell) GetValue() *int {
	return c.val
}

func (c *Cell) GetAllowedVals() []int {
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

func (c *Cell) GetVals() []int {
	return appendIfUniqueMulti(c.row.GetVals(), appendIfUniqueMulti(c.col.GetVals(), c.block.GetVals()...)...)
}

func (c *Cell) GetRow() *Row {
	return c.row
}

func (c *Cell) GetColumn() *Column {
	return c.col
}

func (c *Cell) GetBlock() *Block {
	return c.block
}

func appendIfUniqueMulti(s []int, vals ...int) []int {
	for i := 0; i < len(vals); i++ {
		s = appendIfUnique(s, vals[i])
	}
	return s
}

func appendIfUnique(s []int, val int) []int {
	for i := 0; i < len(s); i++ {
		if val == s[i] {
			return s
		}
	}
	return append(s, val)
}
