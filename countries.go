package countries

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CountryData struct {
	countries map[string]Country
}

type Country struct {
	// ISO 3166-1 alpha-2 code
	Code      string
	Name      string
	Continent string
}

func InstanceCountry() *CountryData {
	instance := &CountryData{}
	instance = instance.readCountriesFile()

	return instance
}

func (data *CountryData) readCountriesFile() *CountryData {
	if len(data.countries) > 0 {
		return data
	}
	path := os.Getenv("GOPATH") + "/src/github.com/RobinToubi/countries/countries.json"
	filePath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(file, &data.countries)
	if err != nil {
		return nil
	}
	return data
}

func (data *CountryData) GetCountryCodeByEnglishName(countryName string) (string, error) {
	for key := range data.countries {
		if data.countries[key].Name == countryName {
			return key, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find country code with English name %s", countryName))
}

func (data *CountryData) GetContinentByCountry(countryCode string) (string, error) {
	if continent, exists := data.countries[countryCode]; exists == false {
		return "", errors.New(fmt.Sprintf("The country %s does not exists in the list", countryCode))
	} else {
		return continent.Continent, nil
	}
}

func (data *CountryData) GetCountryCodesByContinent(continent string) ([]string, error) {
	var codes []string
	for key := range data.countries {
		if data.countries[key].Continent == continent {
			_ = append(codes, key)
		}
	}
	return codes, nil
}
