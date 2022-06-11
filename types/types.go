package types

import "encoding/xml"

type FileType string

type Plural struct {
	XMLName xml.Name      `xml:"plurals" json:"plurals" yaml:"plurals" toml:"plurals" watson:"plurals"`
	Name    string        `xml:"name,attr" json:"name" yaml:"name" toml:"name" watson:"name"`
	Items   []*PluralItem `xml:"item" json:"items" yaml:"items,flow" toml:"items,multiline" watson:"items"`
}

type PluralItem struct {
	XMLName  xml.Name `xml:"item" json:"item" yaml:"item" toml:"item" watson:"item"`
	Quantity string   `xml:"quantity,attr" json:"quantity" yaml:"quantity" toml:"quantity" watson:"quantity"`
	Value    string   `xml:",innerxml" json:"value" yaml:"value" toml:"value" watson:"value"`
}

type Item struct {
	XMLName xml.Name `xml:"item" json:"item" yaml:"item" toml:"item" watson:"item"`
	Value   string   `xml:",innerxml" json:"value" yaml:"value" toml:"value" watson:"value"`
}

type StringArray struct {
	XMLName xml.Name `xml:"string-array" json:"string-array" yaml:"string-array" toml:"string-array" watson:"string-array"`
	Name    string   `xml:"name,attr" json:"name" yaml:"name" toml:"name" watson:"name"`
	Items   []*Item  `xml:"item" json:"items" yaml:"items,flow" toml:"items,multiline" watson:"items"`
}

type String struct {
	XMLName xml.Name `xml:"string" json:"string" yaml:"string" toml:"string" watson:"string"`
	Name    string   `xml:"name,attr" json:"name" yaml:"name" toml:"name" watson:"name"`
	Value   string   `xml:",innerxml" json:"value" yaml:"value" toml:"value" watson:"value"`
}

type Nesting struct {
	XMLName      xml.Name       `xml:"resources" json:"resources" yaml:"resources" toml:"resources" watson:"resources"`
	Strings      []*String      `xml:"string" json:"string" yaml:"string,flow" toml:"string,multiline" watson:"string"`
	StringsArray []*StringArray `xml:"string-array" json:"string-array" yaml:"string-array,flow" toml:"string-array,multiline" watson:"string-array"`
	Plurals      []*Plural      `xml:"plurals" json:"plurals" yaml:"plurals,flow" toml:"plurals,multiline" watson:"plurals"`
}
