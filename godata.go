package godata

import (
	"encoding/json"
	"log"
	"os"
        "fmt"
)

func DataToFile(path string, data ...interface{}) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
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
			log.Println(err)
		}
		_, err = file.WriteString(string(text) + "\n")
		if err != nil {
			log.Println(err)
		}
	}
}
