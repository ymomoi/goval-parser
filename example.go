package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
	"github.com/ymomoi/goval-parser/oval"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s file1 file2 ...\n", os.Args[0])
		return
	}
	for _, f := range os.Args[1:len(os.Args)] {
		fmt.Print(f + " : ")
		oval, err := readOval(f)
		if err != nil {
			fmt.Println(err.Error())
		}
		pp.Println(oval)
	}
}

// readOval : Read OVAL definitions from file
func readOval(file string) (*oval.Root, error) {
	str, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Can't open file: %s", err)
	}
	oval := &oval.Root{}
	err = xml.Unmarshal([]byte(str), oval)
	if err != nil {
		return nil, fmt.Errorf("Can't parse XML: %s", err)
	}
	return oval, nil
}
