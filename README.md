![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Vinetwigs/stres)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Vinetwigs/stres)
![GitHub last commit](https://img.shields.io/github/last-commit/Vinetwigs/stres)
[![stars - stres](https://img.shields.io/github/stars/Vinetwigs/stres?style=social)](https://github.com/Vinetwigs/stres)
[![forks - stres](https://img.shields.io/github/forks/Vinetwigs/stres?style=social)](https://github.com/Vinetwigs/stres)
[![License](https://img.shields.io/badge/License-Apache_License_2.0-orange)](#license)
[![issues - vilmos](https://img.shields.io/github/issues/Vinetwigs/stres)](https://github.com/Vinetwigs/stres/issues)

<h1 align="center">stres - String resources in Go</h1>
<p align="center">
   <i>
        A simple and easy-to-use library to import and export string resources into your Go applications just like you would do in Android Studio.
        Useful to separate business logic from UI texts or to handle string translations. 
   </i>
</p>

## Table of Contents

[TOC]
    

### References

* [Android Studio string resources](https://developer.android.com/guide/topics/resources/string-resource)

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
	XMLName xml.Name      `xml:"plurals"`
	Name    string        `xml:"name,attr"`
	Items   []*PluralItem `xml:"item"`
}
```

[Back to top](#table-of-contents)

#### PluralItem

```go
type PluralItem struct {
	XMLName  xml.Name `xml:"item"`
	Quantity string   `xml:"quantity,attr"`
	Value    string   `xml:",innerxml"`
}
```

[Back to top](#table-of-contents)

#### Item

```go
type Item struct {
	XMLName xml.Name `xml:"item"`
	Value   string   `xml:",innerxml"`
}
```

[Back to top](#table-of-contents)

#### StringArray

```go
type StringArray struct {
	XMLName xml.Name `xml:"string-array"`
	Name    string   `xml:"name,attr"`
	Items   []*Item  `xml:"item"`
}
```

[Back to top](#table-of-contents)

#### String

```go
type String struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",innerxml"`
}
```

[Back to top](#table-of-contents)

#### Nesting

```go
type Nesting struct {
	XMLName      xml.Name       `xml:"resources"`
	Strings      []*String      `xml:"string"`
	StringsArray []*StringArray `xml:"string-array"`
	Plurals      []*Plural      `xml:"plurals"`
}
```

[Back to top](#table-of-contents)

### CreateXMLFile
*Creates strings.xml file in "strings" directory, throws an error otherwise.*

`file, err := stres.CreateXMLFile()`  

[Back to top](#table-of-contents)

### DeleteXMLFile
*Deletes XML file if exists, throws an error otherwise.*

`err := stres.DeleteXMLFile()`

[Back to top](#table-of-contents)

### LoadValues
*Loads values from strings.xml file into internal dictionaries. Needs to be invoked only one time (but before getting strings values).*

`err := stres.LoadValues()`

[Back to top](#table-of-contents)

### NewString
*Adds a new string resource to XML file. Throws an error if the chosen name is already inserted or it is an empty string. Used for programmatic insertion (manual insertion recommended).*

`String, err := stres.NewString("name", "value")`  


| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| name      | string | unique name to identify the string    |   |   |   |   |   |   |   |
| value     | string | string to associate to the given name |

Returns String instance and error.

[Back to top](#table-of-contents)

### NewStringArray
*Adds a new string-array resource to XML file. Throws an error if the chosen name is already inserted or it is an empty string. Used for programmatic insertion (manual insertion recommended).*

`strArr, err := stres.NewStringArray("name", []string{"value1","value2",...})`   

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| name      | string | unique name to identify the string    |   |   |   |   |   |   |   |
| values     | []string | array of strings to associate to the given name |

Returns StringArray instance and error.

[Back to top](#table-of-contents)

### NewQuantityString
*Adds a new quantity string resource to XML file. Throws an error if the chosen name is already inserted or it is an empty string. The function uses only the first 5 values in the array. The first value is assigned to "zero" quantity. The second value is assigned to "one" quantity. The third value is assigned to "two" quantity. The fourth value is assigned to "few" quantity. The fifth value is assigned to "more" quantity. Used for programmatic insertion (manual insertion recommended).*

`qntStr, err := stres.NewQuantityString("name", []string{"zero","one", "two", ...})`   

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| name      | string | unique name to identify the string    |   |   |   |   |   |   |   |
| values     | []string | array of strings for quantities to associate to the given name |

Returns Plural instance and error.

[Back to top](#table-of-contents)

### SetFewThreshold
*Sets the threshold for "few" values in quantity strings.When getting quantity strings values, the function checks if the given count is less OR EQUAL to this value.(default value: 20)*

`stres.SetFewThreshold(25)`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| value      | int | new value for 'few' threshold    |   |   |   |   |   |   |   |

Returns Plural instance and error.

[Back to top](#table-of-contents)

### GetString
*Returns the string resource's value with the given name. If not exists, returns empty string.*

`str := GetString("name")`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| name      | string | unique name given to the corresponding string    |   |   |   |   |   |   |   |

Returns a string.

[Back to top](#table-of-contents)

### GetArrayString
*Returns the string-array resource's values with the given name. If not exists, returns nil.*

`strArr := GetStringArray("name")`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| name      | string | unique name given to the corresponding string-array    |   |   |   |   |   |   |   |

Returns an array of strings.

[Back to top](#table-of-contents)

### GetQuantityString
*Returns the quantity string resource's corresponding string value based on the value of the given count parameter. If the plural is not found, returns an empty string.*

`strArr := GetQuantityString("name", 10)`

| Parameter | Type   | Description                                             |   
|-----------|--------|---------------------------------------|---|---|---|---|---|---|---|
| name      | string | unique name to identify the string    |   |   |   |   |   |   |   |
| count     | int | quantity to fetch the corresponding string |

Returns a string.

[Back to top](#table-of-contents)