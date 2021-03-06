// Code generated by go-enum DO NOT EDIT.
// Version: 0.3.11
// Revision: e477e74e05be2411bb0b218a4637de6905e8cd8c
// Build Date: 2022-02-11T00:05:30Z
// Built By: goreleaser

package enums

import (
	"fmt"
	"strings"
)

const (
	// CityMoscow is a City of type Moscow.
	CityMoscow City = iota
	// CityYekaterinburg is a City of type Yekaterinburg.
	CityYekaterinburg
)

const _CityName = "MoscowYekaterinburg"

var _CityNames = []string{
	_CityName[0:6],
	_CityName[6:19],
}

// CityNames returns a list of possible string values of City.
func CityNames() []string {
	tmp := make([]string, len(_CityNames))
	copy(tmp, _CityNames)
	return tmp
}

var _CityMap = map[City]string{
	CityMoscow:        _CityName[0:6],
	CityYekaterinburg: _CityName[6:19],
}

// String implements the Stringer interface.
func (x City) String() string {
	if str, ok := _CityMap[x]; ok {
		return str
	}
	return fmt.Sprintf("City(%d)", x)
}

var _CityValue = map[string]City{
	_CityName[0:6]:                   CityMoscow,
	strings.ToLower(_CityName[0:6]):  CityMoscow,
	_CityName[6:19]:                  CityYekaterinburg,
	strings.ToLower(_CityName[6:19]): CityYekaterinburg,
}

// ParseCity attempts to convert a string to a City.
func ParseCity(name string) (City, error) {
	if x, ok := _CityValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _CityValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return City(0), fmt.Errorf("%s is not a valid City, try [%s]", name, strings.Join(_CityNames, ", "))
}

func (x City) Ptr() *City {
	return &x
}

// MarshalText implements the text marshaller method.
func (x City) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *City) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseCity(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
