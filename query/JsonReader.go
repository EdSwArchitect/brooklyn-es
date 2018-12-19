package query

import (
	"fmt"
	"io/ioutil"
)

// ReadJSON read the json file
func ReadJSON(name string) string {

	b, err := ioutil.ReadFile(name) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(b) // print the content as 'bytes'

	json := string(b) // convert content to a 'string'

	fmt.Println(json) // print the content as a 'string'

	return json

}
