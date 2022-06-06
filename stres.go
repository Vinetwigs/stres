package stres

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var quantityValues = [...]string{"zero", "one", "two", "few", "many"}

var data []byte
var string_entries map[string]string = make(map[string]string)
var string_array_entries map[string]StringArray = make(map[string]StringArray)
var plural_string_entries map[string]Plural = make(map[string]Plural)

var few_threshold = 20

var (
	ErrorEmptyStringName         error = errors.New("stres: string name can't be empty")
	ErrorEmptyStringArrayName    error = errors.New("stres: string-array name can't be empty")
	ErrorEmptyQuantityStringName error = errors.New("stres: quantity string name can't be empty")

	ErrorDuplicateStringName         error = errors.New("stres: string name already inserted")
	ErrorDuplicateStringArrayName    error = errors.New("stres: string-array name already inserted")
	ErrorDuplicateQuantityStringName error = errors.New("stres: quantity string name already inserted")

	ErrorQuantityStringPluralNotFound error = errors.New("stres: plural not found for the given quantity")

	ErrorQuantityStringEmptyValues error = errors.New("stres: provided empty array to quantity string creationg")
)

type Plural struct {
	XMLName xml.Name      `xml:"plurals"`
	Name    string        `xml:"name,attr"`
	Items   []*PluralItem `xml:"item"`
}

type PluralItem struct {
	XMLName  xml.Name `xml:"item"`
	Quantity string   `xml:"quantity,attr"`
	Value    string   `xml:",innerxml"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Value   string   `xml:",innerxml"`
}

type StringArray struct {
	XMLName xml.Name `xml:"string-array"`
	Name    string   `xml:"name,attr"`
	Items   []*Item  `xml:"item"`
}

type String struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",innerxml"`
}

type Nesting struct {
	XMLName      xml.Name       `xml:"resources"`
	Strings      []*String      `xml:"string"`
	StringsArray []*StringArray `xml:"string-array"`
	Plurals      []*Plural      `xml:"plurals"`
}

/*
	Loads values from strings.xml file into internal dictionaries.
	Needs to be invoked only one time (but before getting strings values).
*/
func LoadValues() error {
	var err error
	n := &Nesting{}

	data, err = readXMLBytes("./strings/strings.xml")
	if err != nil {
		return err
	}

	err = decodeXML(data, &n)
	if err != nil {
		return err
	}

	// Load strings
	for i := 0; i < len(n.Strings); i++ {
		string_entries[n.Strings[i].Name] = n.Strings[i].Value
	}

	// Load string arrays
	for i := 0; i < len(n.StringsArray); i++ {
		string_array_entries[n.StringsArray[i].Name] = *n.StringsArray[i]
	}

	// Load quantity strings
	for i := 0; i < len(n.Plurals); i++ {
		plural_string_entries[n.Plurals[i].Name] = *n.Plurals[i]
	}

	return nil
}

/*
	Creates strings.xml file in "strings" directory, throws an error otherwise.
*/
func CreateXMLFile() (*os.File, error) {
	os.Mkdir("strings", os.ModePerm)

	file, err := os.Create("strings/strings.xml")
	if err != nil {
		return nil, err
	}

	_, err = NewString("name", "value")
	if err != nil {
		return nil, err
	}

	return file, nil
}

/*
	Deletes XML file if exists, throws an error otherwise.
*/
func DeleteXMLFile() error {
	err := os.Remove("strings/strings.xml")
	if err != nil {
		return err
	}

	err = os.Remove("strings")
	if err != nil {
		return err
	}

	return nil
}

/*
	Adds a new string resource to XML file. Throws an error if the chosen name is already inserted or it is an empty string.
*/
func NewString(name, value string) (String, error) {
	if strings.TrimSpace(name) == "" {
		return *new(String), ErrorEmptyStringName
	}

	if isDuplicateString(name) {
		return *new(String), ErrorDuplicateStringName
	}

	string_entries[name] = value

	var err error

	s := String{
		XMLName: xml.Name{},
		Name:    name,
		Value:   value,
	}

	n := &Nesting{}

	data, err = readXMLBytes("./strings/strings.xml")
	if err != nil {
		return *new(String), err
	}

	err = decodeXML(data, &n)
	if err != nil {
		return *new(String), err
	}
	n.Strings = append(n.Strings, &s)

	data, err = encodeXML(n)
	if err != nil {
		return *new(String), err
	}

	err = writeXMLBytes("strings/strings.xml", data)
	if err != nil {
		return *new(String), err
	}

	return s, nil
}

/*
	Adds a new string-array resource to XML file. Throws an error if the chosen name is already inserted or it is an empty string.
*/
func NewStringArray(name string, values []string) (StringArray, error) {
	if strings.TrimSpace(name) == "" {
		return *new(StringArray), ErrorEmptyStringArrayName
	}

	if isDuplicateStringArray(name) {
		return *new(StringArray), ErrorDuplicateStringArrayName
	}

	sa := &StringArray{Name: name}
	for i := 0; i < len(values); i++ {
		item := &Item{
			XMLName: xml.Name{},
			Value:   values[i],
		}
		sa.Items = append(sa.Items, item)
	}

	string_array_entries[name] = *sa

	var err error

	n := &Nesting{}

	data, err = readXMLBytes("./strings/strings.xml")
	if err != nil {
		return *new(StringArray), err
	}

	err = decodeXML(data, &n)
	if err != nil {
		return *new(StringArray), err
	}
	n.StringsArray = append(n.StringsArray, sa)

	data, err = encodeXML(n)
	if err != nil {
		return *new(StringArray), err
	}

	err = writeXMLBytes("strings/strings.xml", data)
	if err != nil {
		return *new(StringArray), err
	}

	return *sa, nil
}

/*
	Adds a new quantity string resource to XML file.
	Throws an error if the chosen name is already inserted or it is an empty string.
	The function uses only the first 5 values in the array.
	The first values is assigned to "zero" quantity.
	The second values is assigned to "one" quantity.
	The third values is assigned to "two" quantity.
	The fourth values is assigned to "few" quantity.
	The fifth values is assigned to "more" quantity.
*/
func NewQuantityString(name string, values []string) (Plural, error) {
	if strings.TrimSpace(name) == "" {
		return *new(Plural), ErrorEmptyStringArrayName
	}

	if len(values) == 0 {
		return *new(Plural), ErrorQuantityStringEmptyValues
	}

	if isDuplicateQuantityString(name) {
		return *new(Plural), ErrorDuplicateQuantityStringName
	}

	pl := &Plural{Name: name}
	for i := 0; i < len(values) && i < 5; i++ {
		item := &PluralItem{
			XMLName:  xml.Name{},
			Quantity: quantityValues[i],
			Value:    values[i],
		}
		pl.Items = append(pl.Items, item)
	}

	plural_string_entries[name] = *pl

	var err error

	n := &Nesting{}

	data, err = readXMLBytes("./strings/strings.xml")
	if err != nil {
		return *new(Plural), err
	}

	err = decodeXML(data, &n)
	if err != nil {
		return *new(Plural), err
	}
	n.Plurals = append(n.Plurals, pl)

	data, err = encodeXML(n)
	if err != nil {
		return *new(Plural), err
	}

	err = writeXMLBytes("strings/strings.xml", data)
	if err != nil {
		return *new(Plural), err
	}

	return *pl, nil
}

/*
	Sets the threshold for "few" values in quantity strings.
	When getting quantity strings values, the function checks if the given count is less OR EQUAL to this value.
	(default value: 20)
*/
func SetFewThreshold(value int) {
	few_threshold = value
}

/*
	Returns the string resource's value with the given name. If not exists, returns empty string.
*/
func GetString(name string) string {
	if name == "" {
		return ""
	}
	if val, ok := string_entries[name]; ok {
		return val
	}
	return ""
}

/*
	Returns the string-array resource's values with the given name. If not exists, returns nil.
*/
func GetArrayString(name string) []string {
	if name == "" {
		return nil
	}

	var arr []string

	if _, ok := string_array_entries[name]; ok {
		for i := 0; i < len(string_array_entries[name].Items); i++ {
			item := string_array_entries[name].Items
			el := item[i].Value
			arr = append(arr, el)
		}
		return arr
	}
	return nil
}

/*
	Returns the quantity string resource's corresponding string value based on the value of the given count parameter.
	If the plural is not found, returns an empty string.
*/
func GetQuantityString(name string, count int) string {
	if name == "" {
		return ""
	}

	val, exists := plural_string_entries[name]

	if !exists {
		return ""
	}

	if count == 0 {
		for i := 0; i < len(val.Items); i++ {
			if val.Items[i].Quantity == quantityValues[0] {
				return val.Items[i].Value
			}
		}
		return ""
	}

	if count == 1 {
		for i := 0; i < len(val.Items); i++ {
			if val.Items[i].Quantity == quantityValues[1] {
				return val.Items[i].Value
			}
		}
		return ""
	}

	if count == 2 {
		for i := 0; i < len(val.Items); i++ {
			if val.Items[i].Quantity == quantityValues[2] {
				return val.Items[i].Value
			}
		}
		return ""
	}

	if count > 2 && count <= few_threshold {
		for i := 0; i < len(val.Items); i++ {
			if val.Items[i].Quantity == quantityValues[3] {
				return val.Items[i].Value
			}
		}
		return ""
	} else {
		for i := 0; i < len(val.Items); i++ {
			if val.Items[i].Quantity == quantityValues[4] {
				return val.Items[i].Value
			}
		}
		return ""
	}
}

func readXMLBytes(path string) ([]byte, error) {
	d, err := ioutil.ReadFile("./strings/strings.xml")
	if err != nil {
		return *new([]byte), err
	}
	return d, nil
}

func decodeXML(data []byte, v interface{}) error {
	if len(data) == 0 {
		v = &Nesting{}
		return nil
	}

	return xml.Unmarshal(data, v)
}

func isDuplicateString(name string) bool {
	if _, ok := string_entries[name]; ok {
		return true
	}
	return false
}

func isDuplicateStringArray(name string) bool {
	if _, ok := string_array_entries[name]; ok {
		return true
	}
	return false
}

func isDuplicateQuantityString(name string) bool {
	if _, ok := plural_string_entries[name]; ok {
		return true
	}
	return false
}

func encodeXML(n *Nesting) ([]byte, error) {
	return xml.MarshalIndent(n, "", "\t")
}

func writeXMLBytes(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}
