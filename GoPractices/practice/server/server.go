package main

import "net/http"

func main() {

	DB = make(map[string][]string)
	http.Serve("")
	client := http.DefaultClient

}

var DB map[string][]string

func DumpData(key, data string) {
	DB[key]
}

func GetDumpData(data string) string {

}
