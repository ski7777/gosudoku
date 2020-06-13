package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/ski7777/gosudoku/internal/grid"
	"github.com/ski7777/gosudoku/internal/loader/online"
	stringloader "github.com/ski7777/gosudoku/internal/loader/string"
	"github.com/ski7777/gosudoku/internal/loadermanager"
	"github.com/ski7777/gosudoku/internal/solvermanager"
)

func main() {
	lm := loadermanager.NewLoadermManager()
	lm.AddLoader(online.Init())
	lm.AddLoader(stringloader.Init())
	loadernames := make([]string, 0)
	for l := range lm.GetLoaders() {
		loadernames = append(loadernames, l)
	}
	parser := argparse.NewParser("GoSudoku", "A complete sudoku solver written in go")
	loader := parser.String("l", "loader", &argparse.Options{Required: true, Help: "Name of the loader. Available: " + strings.Join(loadernames, ", ")})
	loaderhelp := parser.Flag("p", "loader-help", &argparse.Options{Help: "Show help for selected loader"})
	loaderargs := parser.List("a", "loader-arg", &argparse.Options{Help: "Loader arguments. key=value"})
	str := parser.Flag("s", "string", &argparse.Options{Help: "Show input as string"})
	maxprocs := parser.Int("n", "procs", &argparse.Options{Help: "Number of maximum allowed processes. 0 disables complex solving, -1 disables process limit", Default: 50})
	maxsolutions := parser.Int("m", "solutions", &argparse.Options{Help: "Number of maximum showd solutions. 0 disables solving, -1 disables solution limit", Default: -1})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}
	l, ok := lm.GetLoaders()[*loader]
	if !ok {
		log.Fatal(errors.New("Loader " + *loader + " not found!"))
	}
	if *loaderhelp {
		println("Help for loader " + l.GetLongName() + " (" + l.GetShortName() + ")")
		println(l.GetDescription())
		return
	}
	lm.ParseArgs(*loaderargs)
	g := grid.NewGrid()
	if e := lm.Load(*loader, g); e != nil {
		log.Fatal(e)
	}
	log.Println("Mission:")
	if *str {
		log.Println("String:", g.ToString())
	}
	g.Print()
	sm := solvermanager.NewSolverManager(g, *str, *maxprocs, *maxsolutions)
	sm.Solve()
}
