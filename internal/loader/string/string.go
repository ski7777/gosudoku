package string

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ski7777/gosudoku/internal/grid"
)

type String struct {
}

func (l *String) Load(g *grid.Grid, args map[string]string) error {
	data, ok := args["data"]
	if !ok {
		return errors.New("Data argument missing")
	}
	if len(data) != 9*9 {
		return errors.New("Data length mismatch")
	}
	i := 0
	for _, c := range data {
		if v, e := strconv.Atoi(string(c)); e != nil {
			return e
		} else {
			g.Cells()[i%9][(i-i%9)/9].SetValue(v)
		}
		i++
	}
	return nil
}

func (l *String) GetShortName() string { return "string" }

func (l *String) GetLongName() string { return "string-loader" }

func (l *String) GetDescription() string {
	return strings.Join([]string{
		"Loads a sudoku from string",
		"Available arguments:",
		" - data: Sudoku line by line. Empty cells are 0. Required",
	}, "\n")
}

func Init() *String {
	return new(String)
}
