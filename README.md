# godata
## GoData Package
The GoData package provides a function for writing data to a file in JSON format. This package is useful for saving data to disk in a standardized format, making it easy to read and parse by other programs.

## Installation
To use this package, you can simply run the following command:
```sh
go get github.com/ayush-raj8/godata
```

## Usage
To use the GoData package, you will need to import it in your Go program:
```sh
import "github.com/username/godata"
```
You can then use the DataToFile function to write data to a file:
```sh
godata.DataToFile("data.json", myData1, myData2, myData3)
```

The function takes a file path and any number of data parameters of any type. It then opens the file at the specified path, creates it if it does not exist, and writes each data parameter to the file in JSON format.

The function handles several data types differently. Complex numbers are formatted as strings, while structs are formatted as JSON with indentation for readability. All other data types are also formatted as JSON.

## Example
```sh
package main

import (
	gd "github.com/ayush-raj8/godata"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := "Hello, world!"
	num := int16(42)
	pi := float64(3.14159)
	data := map[string]interface{}{
		"key1": "value1",
		"key2": 2,
		"key3": true,
	}
	p := Person{Name: "Alice", Age: 30}
	list := []int{1, 2, 3}
	cmplx := complex(1.0, 2.0)
	ptr := &struct{ Foo string }{Foo: "bar"}
	bo := false

	gd.DataToFile("abc.txt", str, num, pi, data, list, p, ptr, cmplx, bo)
}
```
abc.txt will be created, it does not exists.  It will contain:
```sh
"Hello, world!"
42
3.14159
{
  "key1": "value1",
  "key2": 2,
  "key3": true
}
[
  1,
  2,
  3
]
{
  "name": "Alice",
  "age": 30
}
{
  "Foo": "bar"
}
(1+2i)
false
```
## Contributing
If you find any issues or have suggestions for improvements, please feel free to open an issue or pull request on the GitHub repository. We welcome any and all contributions to make this package better!
