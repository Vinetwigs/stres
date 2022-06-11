package stres

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Vinetwigs/stres/types"
)

var fileType types.FileType
var encDec = types.EncoderDecoder{}

var quantityValues = [...]string{"zero", "one", "two", "few", "many"}

var data []byte
var string_entries map[string]string = make(map[string]string)
var string_array_entries map[string]types.StringArray = make(map[string]types.StringArray)
var plural_string_entries map[string]types.Plural = make(map[string]types.Plural)

var few_threshold = 20

const (
	XML    types.FileType = "xml"
	YAML   types.FileType = "yml"
	JSON   types.FileType = "json"
	TOML   types.FileType = "toml"
	WATSON types.FileType = "watson"
)

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

/*
	Loads values from strings file into internal dictionaries.
	Needs to be invoked only one time (but before getting strings values).
	Takes a FileType parameter to specify strings file format.
*/
func LoadValues(t types.FileType) error {
	SetResourceType(t)

	var err error
	n := &types.Nesting{}

	data, err = readBytes(strings.Join([]string{"./strings/strings.", string(t)}, ""))
	if err != nil {
		return err
	}

	err = encDec.Decode(data, &n)
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
	Used to specify string file extension. If t is a wrong FileType, sets resource type to XML by default.
*/

func SetResourceType(t types.FileType) {
	switch t {
	case "xml":
		xmlED := &types.XMLStrategy{}
		encDec.SetStrategy(xmlED)
		fileType = XML
	case "json":
		jsonED := &types.JSONStrategy{}
		encDec.SetStrategy(jsonED)
		fileType = JSON
	case "yml":
		yamlED := &types.YAMLStrategy{}
		encDec.SetStrategy(yamlED)
		fileType = YAML
	case "toml":
		tomlED := &types.TOMLStrategy{}
		encDec.SetStrategy(tomlED)
		fileType = TOML
	case "watson":
		watsonED := &types.WatsonStrategy{}
		encDec.SetStrategy(watsonED)
		fileType = WATSON
	default:
		xmlED := &types.XMLStrategy{}
		encDec.SetStrategy(xmlED)
		fileType = XML
	}
}

/*
	Creates strings resource file in "strings" directory, throws an error otherwise.
	Takes a FileType parameter to specify strings file format.
*/
func CreateResourceFile(t types.FileType) (*os.File, error) {

	fileType = t

	SetResourceType(t)

	os.Mkdir("strings", os.ModePerm)

	file, err := os.Create("strings/strings." + string(fileType))
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
	Deletes resource file if exists, throws an error otherwise.
	Uses setted resource file extension.
*/
func DeleteResourceFile() error {
	err := os.Remove("strings/strings." + string(fileType))
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
	Adds a new string resource to resource file. Throws an error if the chosen name is already inserted or it is an empty string.
*/
func NewString(name, value string) (types.String, error) {
	if strings.TrimSpace(name) == "" {
		return *new(types.String), ErrorEmptyStringName
	}

	if isDuplicateString(name) {
		return *new(types.String), ErrorDuplicateStringName
	}

	string_entries[name] = value

	var err error

	s := types.String{
		Name:  name,
		Value: value,
	}

	n := &types.Nesting{}

	data, err = readBytes(strings.Join([]string{"./strings/strings.", string(fileType)}, ""))
	if err != nil {
		return *new(types.String), err
	}

	err = encDec.Decode(data, &n)
	if err != nil {
		return *new(types.String), err
	}
	n.Strings = append(n.Strings, &s)

	data, err = encDec.Encode(n)

	for i := 0; i < len(data); i++ {
		fmt.Printf("%+v, ", data[i])
		if i%10 == 0 {
			println()
		}
	}
	println()

	if err != nil {
		return *new(types.String), err
	}

	err = writeBytes(strings.Join([]string{"./strings/strings.", string(fileType)}, ""), data)
	if err != nil {
		return *new(types.String), err
	}

	return s, nil
}

/*
	Adds a new string-array resource to resource file. Throws an error if the chosen name is already inserted or it is an empty string.
*/
func NewStringArray(name string, values []string) (types.StringArray, error) {
	if strings.TrimSpace(name) == "" {
		return *new(types.StringArray), ErrorEmptyStringArrayName
	}

	if isDuplicateStringArray(name) {
		return *new(types.StringArray), ErrorDuplicateStringArrayName
	}

	sa := &types.StringArray{Name: name}
	for i := 0; i < len(values); i++ {
		item := &types.Item{
			Value: values[i],
		}
		sa.Items = append(sa.Items, item)
	}

	string_array_entries[name] = *sa

	var err error

	n := &types.Nesting{}

	data, err = readBytes("./strings/strings." + string(fileType))
	if err != nil {
		return *new(types.StringArray), err
	}

	err = encDec.Decode(data, &n)
	if err != nil {
		return *new(types.StringArray), err
	}
	n.StringsArray = append(n.StringsArray, sa)

	data, err = encDec.Encode(n)
	if err != nil {
		return *new(types.StringArray), err
	}

	err = writeBytes("strings/strings."+string(fileType), data)
	if err != nil {
		return *new(types.StringArray), err
	}

	return *sa, nil
}

/*
	Adds a new quantity string resource to resource file.
	Throws an error if the chosen name is already inserted or it is an empty string.
	The function uses only the first 5 values in the array.
	The first values is assigned to "zero" quantity.
	The second values is assigned to "one" quantity.
	The third values is assigned to "two" quantity.
	The fourth values is assigned to "few" quantity.
	The fifth values is assigned to "more" quantity.
*/
func NewQuantityString(name string, values []string) (types.Plural, error) {
	if strings.TrimSpace(name) == "" {
		return *new(types.Plural), ErrorEmptyStringArrayName
	}

	if len(values) == 0 {
		return *new(types.Plural), ErrorQuantityStringEmptyValues
	}

	if isDuplicateQuantityString(name) {
		return *new(types.Plural), ErrorDuplicateQuantityStringName
	}

	pl := &types.Plural{Name: name}
	for i := 0; i < len(values) && i < 5; i++ {
		item := &types.PluralItem{
			Quantity: quantityValues[i],
			Value:    values[i],
		}
		pl.Items = append(pl.Items, item)
	}

	plural_string_entries[name] = *pl

	var err error

	n := &types.Nesting{}

	data, err = readBytes("./strings/strings." + string(fileType))
	if err != nil {
		return *new(types.Plural), err
	}

	err = encDec.Decode(data, &n)
	if err != nil {
		return *new(types.Plural), err
	}
	n.Plurals = append(n.Plurals, pl)

	data, err = encDec.Encode(n)
	if err != nil {
		return *new(types.Plural), err
	}

	err = writeBytes("strings/strings."+string(fileType), data)
	if err != nil {
		return *new(types.Plural), err
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

func readBytes(path string) ([]byte, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return *new([]byte), err
	}
	return d, nil
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

func writeBytes(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}
