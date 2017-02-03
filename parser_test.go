package oval

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestReading(t *testing.T) {
	file := "testfile/cisco-sa-20160928-msdp_613.xml"
	oval := &OVALDefinitions{}
	err := ReadOval(file, oval)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
	pp.Println(oval)

	file = "testfile/cisco-sa-20160323-smi_583.xml"
	oval = &OVALDefinitions{}
	err = ReadOval(file, oval)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
	pp.Println(oval)
}

/*
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
}
*/
