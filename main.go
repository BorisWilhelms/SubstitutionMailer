// GoTest project main.go
package main

import (
	"encoding/base64"
	"fmt"
	_ "github.com/BorisWilhelms/GoUtils/Unmarshal"
	"github.com/BorisWilhelms/SubstitutionMailer/Substitution"
	"io/ioutil"
)

func main() {
	source, err := ioutil.ReadFile("source.htm")
	if err != nil {
		fmt.Println("Source could not be read... :%s", err.Error())
	}
	s, err := Substitution.Parse(string(source), "6b")
	if err != nil {
		fmt.Println("Source could not be parsed... :%s", err.Error())
	}
	fmt.Printf("%#v\n", s)
}
