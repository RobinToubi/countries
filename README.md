# countries
Go package to retrieve country data by the ISO 3166-1 alpha-2 code.

-----

[![Build & Tests](https://github.com/RobinToubi/countries/actions/workflows/go.yml/badge.svg)](https://github.com/RobinToubi/countries/actions/workflows/go.yml)

## Install

```
go get github.com/RobinToubi/countries
```

## Use

#### Continent name by country code

```
func GetContinentByCountry(countryCode string) string, error
```

#### Continent name To Country code list

```
func GetCountryCodesByContinent(continent string) []string
```

## Simple example

```go
package main

import (
	"fmt"
	"github.com/RobinToubi/countries"
)

func main() {
	// Instantiate the service
	countryService := countries.InstanceCountry()
	
	// Retrieve a continent by the country code
	continent, _ := countryService.GetContinentByCountry("HK")
	fmt.Println(continent) // "Asia"
	
	// Retrieve country codes by a continent name
	cCodes, _ := countryService.GetCountryCodesByContinent("Asia")
	fmt.Println(cCodes) // [AF AM AZ BH BD BT BN KH CN CX CC CY GE HK IN ID IR IQ IL JP JO KZ KP KR KW KG LA LB MY MV MN MM NP OM PK PH QA RU SA SG LK SY TW TJ TH TR TM AE UZ VN YE]
}
```
