package countries

import (
	"fmt"
	"testing"
)

var countryData *CountryData

func init() {
	countryData = InstanceCountry()
}

func TestData_GetCountryCodeByEnglishName(t *testing.T) {
	expected := "US"
	given, err := countryData.GetCountryCodeByEnglishName("United States")
	if given != expected && err != nil {
		fmt.Println(err)
		t.Fatalf(`TestData_GetCountryCodeByEnglishName("United States") = %s, want "%s", error`, given, expected)
	}
}

func TestGetContinentByCountryValid(t *testing.T) {
	expected := "Asia"
	given, err := countryData.GetContinentByCountry("CN")
	if expected != given && err != nil {
		t.Fatalf(`GetContinentByCountry("CN") = %s, want "%s", error`, given, expected)
	}
}

func TestGetContinentByCountryError(t *testing.T) {
	expected := ""
	given, err := countryData.GetContinentByCountry("")
	if expected != given && err != nil {
		t.Fatalf(`GetContinentByCountry("") = %s, want "%s", error`, given, expected)
	}
}

func TestGetCountriesByContinent(t *testing.T) {
	expected := []string{"AS",
		"AU",
		"CK",
		"FJ",
		"PF",
		"GU",
		"KI",
		"MH",
		"FM",
		"NR",
		"NC",
		"NZ",
		"NU",
		"NF",
		"MP",
		"PW",
		"PG",
		"PN",
		"SB",
		"TK",
		"TO",
		"TV",
		"VU",
		"WF",
		"WS"}
	given, err := countryData.GetCountryCodesByContinent("Oceania")
	if len(expected) != len(given) && err != nil {
		t.Fatalf(`GetCountriesByContinent("Oceania") = %d, want "%d", error`, len(given), len(expected))
	}
}

func TestGetCountriesByContinentEmpty(t *testing.T) {
	expected := make([]string, 0)
	given, err := countryData.GetCountryCodesByContinent("")
	if len(expected) != len(given) && err != nil {
		t.Fatalf(`GetCountriesByContinent("Oceania") = %d, want "%d", error`, len(given), len(expected))
	}
}
