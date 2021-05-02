package simple

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GenerateRPC(filename string) {
	schema, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		return
	}
	lines := strings.Split(string(schema), "\n")
	fmt.Println(len(lines))
	service := false
	name := ""
	for _, line := range lines {
		if strings.Contains(line, "service ") {
			tokens := strings.Split(line, " ")
			name = tokens[1]
			service = true
		}
		if service {
			fmt.Println(name, line)
		}
	}
}
