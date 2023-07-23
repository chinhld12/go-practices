package sortarray

import (
	"sort"

	"github.com/chinhld12/assignment-02/util"
)

func sortNumbers(input *[]float64) {
	sort.Slice(*input, func(i, j int) bool {
		return (*input)[i] < (*input)[j]
	})
}

func sortStrings(input *[]string) {
	sort.Strings(*input)
}

// HandleSortStrings sorts a slice of strings
func HandleSortStrings(input []string) []string {
	allStrings := util.FilterOnlyStrings(input)
	sortStrings(&allStrings)
	return allStrings
}

// HandleSortFloats sorts a slice of floats
func HandleSortFloats(input []string) []string {
	allFloatInt := util.FilterOnlyFloats(input)
	sortNumbers(&allFloatInt)
	return util.TransformFloatToString(allFloatInt)
}

// HandleSortMixes sorts a slice of mixed types
func HandleSortMixes(input []string) []string {
	allFloatInt := util.FilterOnlyFloats(input)
	allStrings := util.FilterOnlyStrings(input)

	sortStrings(&allStrings)
	sortNumbers(&allFloatInt)

	stringifiedFloat := util.TransformFloatToString(allFloatInt)
	return append(stringifiedFloat, allStrings...)
}
