package simple

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GenerateModelRPC(outdir, filename string) {
	schema, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		return
	}
	lines := strings.Split(string(schema), "\n")
	fmt.Println(len(lines))
}
func GenerateFuncRPC(outdir, filename string) {
	schema, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		return
	}
	lines := strings.Split(string(schema), "\n")
	fmt.Println(len(lines))
}
