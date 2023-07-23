package cli

import (
	"flag"
	"fmt"
	"strings"
	"github.com/chinhld12/assignment-02/internal/sortarray"
)

var intF, floatF, mixF, stringF bool

func init() {
	flag.BoolVar(&stringF, "string", false, "Sort array of strings")
	flag.BoolVar(&intF, "int", false, "Sort array of int")
	flag.BoolVar(&floatF, "float", false, "Sort array of float")
	flag.BoolVar(&mixF, "mix", false, "Sort array of mix of strings and numbers")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	flag.Parse()
	args := flag.Args()

	switch {
	case stringF:
		res := sortarray.HandleSortStrings(args)
		fmt.Println(strings.Join(res, " "))
	case floatF, intF:
		res := sortarray.HandleSortFloats(args)
		fmt.Println(strings.Join(res, " "))
	case mixF:
		res := sortarray.HandleSortMixes(args)
		fmt.Println(strings.Join(res, " "))
	default:
		panic("Please input correctly flag")
	}
}
