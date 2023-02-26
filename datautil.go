package datautil

package main

import (
	"encoding/json"
	"os"
)


func DataToFile(path string, data ...interface{}) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, d := range data {
		var text []byte
		switch v := d.(type) {
		case complex64, complex128:
			text = []byte(fmt.Sprintf("%v", d))
		case *struct{}:
			text, err = json.MarshalIndent(v, "", "  ")
		default:
			text, err = json.MarshalIndent(d, "", "  ")
		}
		if err != nil {
			fmt.Println(err)
		}
		_, err = file.WriteString(string(text) + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}
