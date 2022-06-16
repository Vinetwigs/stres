[![Go Report Card](https://goreportcard.com/badge/github.com/Vinetwigs/stres)](https://goreportcard.com/report/github.com/Vinetwigs/stres)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Vinetwigs/stres)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Vinetwigs/stres)
![GitHub last commit](https://img.shields.io/github/last-commit/Vinetwigs/stres)
[![stars - stres](https://img.shields.io/github/stars/Vinetwigs/stres?style=social)](https://github.com/Vinetwigs/stres)
[![forks - stres](https://img.shields.io/github/forks/Vinetwigs/stres?style=social)](https://github.com/Vinetwigs/stres)
[![License](https://img.shields.io/badge/License-Apache_License_2.0-orange)](#license)
[![issues - vilmos](https://img.shields.io/github/issues/Vinetwigs/stres)](https://github.com/Vinetwigs/stres/issues)

<h1 align="center">stres - Android Studio string resources in Go</h1>
<p align="center">
   <i>
        A simple and easy-to-use library to import and export string resources into your Go applications just like you would do in Android Studio.
        Useful to separate business logic from UI texts or to handle string translations. 
   </i>
</p>

## Table of contents

  * [References](#references)
  * [Prerequisites](#prerequisites)
  * [Installing](#installing)
    + [Install module](#install-module)
    + [Import in your project](#import-in-your-project)
- [Documentation](#documentation)
  * [Types](#types)
    + [Plural](#plural)
    + [PluralItem](#pluralitem)
    + [Item](#item)
    + [StringArray](#stringarray)
    + [String](#string)
    + [Nesting](#nesting)
  * [CreateResourceFile](#createresourcefile)
  * [DeleteResourceFile](#deleteresourcefile)
  * [LoadValues](#loadvalues)
  * [SetResourceType](#setresourcetype)
  * [NewString](#newstring)
  * [NewStringArray](#newstringarray)
  * [NewQuantityString](#newquantitystring)
  * [SetFewThreshold](#setfewthreshold)
  * [GetString](#getstring)
  * [GetArrayString](#getarraystring)
  * [GetQuantityString](#getquantitystring)
- [Contributors](#contributors)


### References

* [Android Studio string resources](https://developer.android.com/guide/topics/resources/string-resource)
* [CHANGELOG](./CHANGELOG.md)

[Back to top](#table-of-contents)

### Prerequisites

* Make sure you have at least installed `Go v1.17`.

[Back to top](#table-of-contents)

### Installing

#### Install module
```
go get github.com/Vinetwigs/stres
```

[Back to top](#table-of-contents)

#### Import in your project
```
import("github.com/Vinetwigs/stres")
```
[Back to top](#table-of-contents)

## Documentation

### Types

#### Plural

```go
type Plural struct {
	XMLName xml.Name      `xml:"plurals" json:"plurals" yaml:"plurals" toml:"plurals" watson:"plurals" msgpack:"plurals"`
	Name    string        `xml:"name,attr" json:"name" yaml:"name" toml:"name" watson:"name" msgpack:"name"`
	Items   []*PluralItem `xml:"item" json:"items" yaml:"items,flow" toml:"items,multiline" watson:"items" msgpack:"items,as_array"`
}
```

[Back to top](#table-of-contents)

#### PluralItem

```go
type PluralItem struct {
	XMLName  xml.Name `xml:"item" json:"item" yaml:"item" toml:"item" watson:"item" msgpack:"item"`
	Quantity string   `xml:"quantity,attr" json:"quantity" yaml:"quantity" toml:"quantity" watson:"quantity" msgpack:"quantity"`
	Value    string   `xml:",innerxml" json:"value" yaml:"value" toml:"value" watson:"value" msgpack:"value"`
}
```

[Back to top](#table-of-contents)

#### Item

```go
type Item struct {
	XMLName xml.Name `xml:"item" json:"item" yaml:"item" toml:"item" watson:"item" msgpack:"item"`
	Value   string   `xml:",innerxml" json:"value" yaml:"value" toml:"value" watson:"value" msgpack:"value"`
}
```

[Back to top](#table-of-contents)

#### StringArray

```go
type StringArray struct {
	XMLName xml.Name `xml:"string-array" json:"string-array" yaml:"string-array" toml:"string-array" watson:"string-array" msgpack:"string-array"`
	Name    string   `xml:"name,attr" json:"name" yaml:"name" toml:"name" watson:"name" msgpack:"name"`
	Items   []*Item  `xml:"item" json:"items" yaml:"items,flow" toml:"items,multiline" watson:"items" msgpack:"items,as_array" `
}
```

[Back to top](#table-of-contents)

#### String

```go
type String struct {
	XMLName xml.Name `xml:"string" json:"string" yaml:"string" toml:"string" watson:"string" msgpack:"string"`
	Name    string   `xml:"name,attr" json:"name" yaml:"name" toml:"name" watson:"name" msgpack:"name"`
	Value   string   `xml:",innerxml" json:"value" yaml:"value" toml:"value" watson:"value" msgpack:"value"`
}
```

[Back to top](#table-of-contents)

#### Nesting

```go
type Nesting struct {
	XMLName      xml.Name       `xml:"resources" json:"resources" yaml:"resources" toml:"resources" watson:"resources" msgpack:"resources"`
	Strings      []*String      `xml:"string" json:"string" yaml:"string,flow" toml:"string,multiline" watson:"string" msgpack:"string,as_array"`
	StringsArray []*StringArray `xml:"string-array" json:"string-array" yaml:"string-array,flow" toml:"string-array,multiline" watson:"string-array" msgpack:"string-array,as_array"`
	Plurals      []*Plural      `xml:"plurals" json:"plurals" yaml:"plurals,flow" toml:"plurals,multiline" watson:"plurals" msgpack:"plurals,as_array"`
}
```

[Back to top](#table-of-contents)

### CreateResourceFile
*Creates strings resource file in "strings" directory, throws an error otherwise. Takes a FileType parameter to specify strings file format.*

`file, err := stres.CreateXMLFile()`

| Parameter | Type   | Description                           |   
|-----------|--------|---------------------------------------|
| t      | types.FileType | enum value to specify file format    |

[Back to top](#table-of-contents)

### DeleteResourceFile
*Deletes resource file if exists, throws an error otherwise. Uses setted resource file extension.*

`err := stres.DeleteXMLFile()`

[Back to top](#table-of-contents)

### LoadValues
*Loads values from strings file into internal dictionaries. Needs to be invoked only one time (but before getting strings values).Takes a FileType parameter to specify strings file format.*

`err := stres.LoadValues()`

[Back to top](#table-of-contents)

| Parameter | Type   | Description                           |   
|-----------|--------|---------------------------------------|
| t      | types.FileType | enum value to specify file format    |

### SetResourceType
*Used to specify string file extension. If t is a wrong FileType, sets resource type to XML by default.*

`stres.SetResourceType(stres.WATSON)`

| Parameter | Type   | Description                           |   
|-----------|--------|---------------------------------------|
| t      | types.FileType | enum value to specify file format    |

[Back to top](#table-of-contents)

### NewString
*Adds a new string resource to resource file. Throws an error if the chosen name is already inserted or it is an empty string. Used for programmatic insertion (manual insertion recommended).*

`String, err := stres.NewString("name", "value")`  


| Parameter | Type   | Description                           |   
|-----------|--------|---------------------------------------|
| name      | string | unique name to identify the string    |
| value     | string | string to associate to the given name |

Returns String instance and error.

[Back to top](#table-of-contents)

### NewStringArray
*Adds a new string-array resource to resource file. Throws an error if the chosen name is already inserted or it is an empty string. Used for programmatic insertion (manual insertion recommended).*

`strArr, err := stres.NewStringArray("name", []string{"value1","value2",...})`   

| Parameter | Type   | Description                           |   
|-----------|--------|---------------------------------------|
| name      | string | unique name to identify the string    |
| values    | []string | array of strings to associate to the given name |

Returns StringArray instance and error.

[Back to top](#table-of-contents)

### NewQuantityString
*Adds a new quantity string resource to resource file. Throws an error if the chosen name is already inserted or it is an empty string. The function uses only the first 5 values in the array. The first value is assigned to "zero" quantity. The second value is assigned to "one" quantity. The third value is assigned to "two" quantity. The fourth value is assigned to "few" quantity. The fifth value is assigned to "more" quantity. Used for programmatic insertion (manual insertion recommended).*

`qntStr, err := stres.NewQuantityString("name", []string{"zero","one", "two", ...})`   

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|
| name      | string | unique name to identify the string    |
| values     | []string | array of strings for quantities to associate to the given name |

Returns Plural instance and error.

[Back to top](#table-of-contents)

### SetFewThreshold
*Sets the threshold for "few" values in quantity strings.When getting quantity strings values, the function checks if the given count is less OR EQUAL to this value.(default value: 20)*

`stres.SetFewThreshold(25)`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|
| value      | int | new value for 'few' threshold    |   

Returns Plural instance and error.

[Back to top](#table-of-contents)

### GetString
*Returns the string resource's value with the given name. If not exists, returns empty string.*

`str := GetString("name")`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|
| name      | string | unique name given to the corresponding string    |

Returns a string.

[Back to top](#table-of-contents)

### GetArrayString
*Returns the string-array resource's values with the given name. If not exists, returns nil.*

`strArr := GetStringArray("name")`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|
| name      | string | unique name given to the corresponding string-array    |

Returns an array of strings.

[Back to top](#table-of-contents)

### GetQuantityString
*Returns the quantity string resource's corresponding string value based on the value of the given count parameter. If the plural is not found, returns an empty string.*

`strArr := GetQuantityString("name", 10)`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|
| name      | string | unique name to identify the string    |
| count     | int | quantity to fetch the corresponding string |

Returns a string.

[Back to top](#table-of-contents)

## Contributors

<a href="https://github.com/Vinetwigs/stres/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Vinetwigs/stres" />
</a>
