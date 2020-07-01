package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const help = `
This beauty help you if you forgot something. 
- what - lists all your available forgotten things
- help - you're reading this nao

`

func main() {
	action := os.Args[1]

	if len(os.Args) == 1 {
		fmt.Printf("Literally don't know what to do\n")
		os.Exit(1)
	}

	switch action {
	case "help":
		fmt.Printf(help)
		os.Exit(1)
	case "what":
		files, err := findAvailable()
		if err != nil {
			fmt.Printf("something went really really bad... %v\n", err)
		}
		fmt.Printf("Available files:\n")
		for _, file := range files {
			fmt.Printf("%v\t", file)
		}
	default:
		err := readFile(action)
		if err != nil {
			fmt.Printf("something went really really bad... %v\n", err)
			os.Exit(1)
		}
	}
}

func findAvailable() ([]string, error) {
	var files []string
	root := "./forgot"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func readFile(name string) error {
	filePath := fmt.Sprintf("./forgot/%v.md", name)
	reader, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fmt.Println(string(reader))

	return nil
}
