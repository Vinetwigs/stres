package stres

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
)

var data []byte

type String struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",innerxml"`
}

type Nesting struct {
	XMLName xml.Name  `xml:"resources"`
	Strings []*String `xml:"string"`
}

// Creates strings.xml file in "strings" directory, throws an error otherwise.
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

// Deletes XML file if exists, throws an error otherwise
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

// Adds a new string resource to XML file. Throws an error if the chosen name is already inserted.
func NewString(name, value string) (String, error) {
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

	if isDuplicate(n, name) {
		return *new(String), errors.New("string name already inserted")
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

// Returns the string resource's value with the given name. If not exists, returns empty string.
func GetString(name string) string {
	var err error

	n := &Nesting{}

	data, err = ioutil.ReadFile("./strings/strings.xml")
	if err != nil {
		return ""
	}
	xml.Unmarshal(data, &n)

	for i := 0; i < len(n.Strings); i++ {
		if n.Strings[i].Name == name {
			return n.Strings[i].Value
		}
	}

	return ""
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

func isDuplicate(n *Nesting, name string) bool {
	for i := 0; i < len(n.Strings); i++ {
		if n.Strings[i].Name == name {
			return true
		}
	}
	return false
}

func encodeXML(n *Nesting) ([]byte, error) {
	return xml.MarshalIndent(n, "", "\t")
}

func writeXMLBytes(path string, data []byte) error {
	return os.WriteFile(path, data, 0666)
}
