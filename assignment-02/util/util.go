package util

import "strconv"

// ParseFloat parses a string into a float64
func ParseFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

// FilterOnlyFloats filters a slice of strings and returns a slice of float64
func FilterOnlyFloats(input []string) []float64 {
	var output []float64
	for _, v := range input {
		if f, err := ParseFloat(v); err == nil {
			output = append(output, f)
		}
	}
	return output
}

// FilterOnlyStrings filters a slice of strings and returns a slice of strings
func FilterOnlyStrings(input []string) []string {
	var output []string
	for _, v := range input {
		if _, err := ParseFloat(v); err != nil {
			output = append(output, v)
		}
	}
	return output
}

// TransformFloatToString transforms a slice of float64 into a slice of strings
func TransformFloatToString(input []float64) []string {
	var output []string
	for _, v := range input {
		output = append(output, strconv.FormatFloat(v, 'f', -1, 64))
	}
	return output
}
