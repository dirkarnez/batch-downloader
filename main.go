package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
)

func main() {
	file, err := os.Open("input.json")
	checkError(err)
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	checkError(err)

	jsonparser.ArrayEach(buf.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		url, err := jsonparser.GetString(value, "url")
		checkError(err)

		name, err := jsonparser.GetString(value, "name")
		checkError(err)

		download(url, name)
		checkError(err)
	})
}

func download(url, filePath string) {
	// Get the data
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	checkError(err)
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
