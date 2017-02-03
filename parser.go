package oval

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// ReadOval : Read OVAL definitions from file
func ReadOval(file string, oval *OVALDefinitions) error {
	str, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Can't open file: %s", err)
	}

	err = xml.Unmarshal([]byte(str), &oval)
	if err != nil {
		return fmt.Errorf("Can't parse XML: %s", err)
	}
	return nil
}
