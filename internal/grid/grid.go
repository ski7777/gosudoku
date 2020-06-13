package grid

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Grid struct {
	cells  [9][9]*Cell
	rows   [9]*Row
	cols   [9]*Column
	blocks [3][3]*Block
}

func (g *Grid) Cells() [9][9]*Cell {
	return g.cells
}

func (g *Grid) Rows() [9]*Row {
	return g.rows
}

func (g *Grid) Cols() [9]*Column {
	return g.cols
}

func (g *Grid) Blocks() [3][3]*Block {
	return g.blocks
}

func (g *Grid) CheckSolved() bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if g.cells[x][y].GetValue() == nil {
				return false
			}
		}
	}
	return true
}

func (g *Grid) ToString() string {
	str := ""
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if v := g.cells[y][x].GetValue(); v == nil {
				str += "0"
			} else {
				str += strconv.Itoa(*v)
			}
		}
	}
	return str
}

func (g *Grid) Print() {
	data := make([][]string, 9)
	for x := 0; x < 9; x++ {
		data[x] = make([]string, 9)
	}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if v := g.cells[x][y].GetValue(); v == nil {
				data[y][x] = " "
			} else {
				data[y][x] = strconv.Itoa(*v)
			}
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func (g *Grid) Clone() *Grid {
	ng := NewGrid()
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if v := g.cells[x][y].GetValue(); v != nil {
				ng.cells[x][y].SetValue(*v)
			}
		}
	}
	return ng
}

func (g *Grid) Equals(o *Grid) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			v1 := g.cells[x][y].GetValue()
			v2 := o.cells[x][y].GetValue()
			if v1 == nil {
				if v2 != nil {
					return false
				}
			} else {
				if v2 == nil {
					return false
				}
				if *v1 != *v2 {
					return false
				}
			}
		}
	}
	return true
}

func NewGrid() *Grid {
	g := new(Grid)
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			g.cells[x][y] = new(Cell)
		}
	}
	var rcells [9]*Cell
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			rcells[x] = g.cells[x][y]
		}
		g.rows[y] = newRow(rcells)
	}
	for x := 0; x < 9; x++ {
		g.cols[x] = newCol(g.cells[x])
	}
	var bcells [3][3]*Cell
	for bx := 0; bx < 3; bx++ {
		for by := 0; by < 3; by++ {
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					bcells[x][y] = g.cells[bx*3+x][by*3+y]
				}
			}
			g.blocks[bx][by] = newBlock(bcells)
		}
	}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			c := g.cells[x][y]
			c.row = g.rows[y]
			c.col = g.cols[x]
			c.block = g.blocks[(x-x%3)/3][(y-y%3)/3]
		}
	}
	return g
}
