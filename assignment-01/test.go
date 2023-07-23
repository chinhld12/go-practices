package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const eastern string = "VN, KR, JP, IR, IQ, SA, CN"

const western string = "DE, FR, GB, US, MX, CA, AU, NZ, PG"

const errorColor string = "\033[1;31m%s\033[0m"

func makeError(text string) error {
	return errors.New(text)
}

func printError(err error) {
	fmt.Printf(errorColor, "Error: "+err.Error())
}

func getIsHasCountryCode[T string](list []T, country T) bool {
	set := make(map[T]struct{}, len(list))
	for _, s := range list {
		set[s] = struct{}{}
	}

	_, ok := set[country]
	return ok
}

func getIsEasternCountries(country string) bool {
	return getIsHasCountryCode(strings.Split(eastern, ", "), country)
}

func getIsWesternCountries(country string) bool {
	return getIsHasCountryCode(strings.Split(western, ", "), country)
}

func orderTheName(args []string) (string, error) {
	firstName := args[0]
	lastName := args[1]
	middleName := ""
	country := ""
	lengthAgrs := len(args)

	if lengthAgrs == 3 {
		country = args[2]
	} else {
		var midLength = lengthAgrs - 1
		middleName = strings.Join(args[2:midLength], " ")
		country = args[midLength]
	}

	if len(country) != 2 {
		return "", makeError("Country code is not valid")
	}

	switch {
	case getIsEasternCountries(country):
		return strings.ReplaceAll(fmt.Sprintf("%s %s %s", lastName, middleName, firstName), "  ", " "), nil
	case getIsWesternCountries(country):
		return strings.ReplaceAll(fmt.Sprintf("%s %s %s", firstName, lastName, middleName), "  ", " "), nil
	default:
		return "", makeError("Country is not matched")
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 3 {
		printError(makeError("Not enough arguments\nOnly accept 3 or 4 arguments"))
		return
	}

	name, err := orderTheName(argsWithoutProg)
	if err != nil {
		printError(err)
		return
	}
	fmt.Println("Output:", name)
}
