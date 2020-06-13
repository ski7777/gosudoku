package online

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	stringloader "github.com/ski7777/gosudoku/internal/loader/string"

	"github.com/ski7777/gosudoku/package/grid"
)

const url = "https://sudoku.com/api/getLevel/"

type Online struct {
}

type Response struct {
	Answer string        `json:"answer"`
	Desc   []interface{} `json:"desc"`
}

func (l *Online) Load(g *grid.Grid, args map[string]string) error {
	level, ok := args["level"]
	if !ok {
		return errors.New("Level argument missing")
	}
	if level != "easy" && level != "medium" && level != "hard" && level != "expert" {
		return errors.New("Level argument invalid")
	}
	r, err := (&http.Client{Timeout: 10 * time.Second}).Get(url + level)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	target := &Response{}
	if e := json.NewDecoder(r.Body).Decode(target); e != nil {
		return e
	}
	if target.Answer != "success" {
		return errors.New("Server error")
	}
	data, ok := target.Desc[0].(string)
	if !ok {
		return errors.New("Data mismatch")
	}
	return stringloader.Init().Load(g, map[string]string{"data": data})
}

func (l *Online) GetShortName() string { return "online" }

func (l *Online) GetLongName() string { return "sudoku.com-loader" }

func (l *Online) GetDescription() string {
	return strings.Join([]string{
		"Loads a sudoku from sudoku.com",
		"Available arguments:",
		" - level: One of: easy, medium, hard, expert. Required",
	}, "\n")
}

func Init() *Online {
	return new(Online)
}
