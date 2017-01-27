package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
	"github.com/ymomoi/goval-parser/oval"
)

func readOval(file string) error {
	str, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Can't open file: %s", err)
	}

	val := &oval.OVALDefinitions{}
	err = xml.Unmarshal([]byte(str), &val)
	if err != nil {
		return fmt.Errorf("Can't parse XML: %s", err)
	}
	pp.Println(val)
	return nil
}

func main() {
	if len(os.Args) > 1 {
		for _, f := range os.Args[1:len(os.Args)] {
			fmt.Print(f + " : ")
			err := readOval(f)
			if err != nil {
				fmt.Print(err)
			} else {
				fmt.Println("success")
			}
		}
	} else {
		fmt.Printf("usage: %s file1 file2 ...\n", os.Args[0])
	}

	/*
		nicoXML := Nicovideo{Thumb{"", ""}}
		err := xml.Unmarshal([]byte(str), &nicoXML)

		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(nicoXML)
		fmt.Println("title: " + nicoXML.Thumb.Title)
		fmt.Println("length: " + nicoXML.Thumb.Length)
	*/
}

/*
type Nicovideo struct {
	Thumb Thumb `xml:"thumb"`
}

type Thumb struct {
	Title  string `xml:"title"`
	Length string `xml:"length"`
}
*/
