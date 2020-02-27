package main

import (
	"paraglider/glider/yamlparser"
	"fmt"
)

func main() {
	c, _ := yamlparser.ParseYaml("../sketch.yaml")

	fmt.Println("+v\n", *c)
	c1 := *c

	fmt.Println("\n")
	fmt.Println(c1.Backends[0].Name)
	fmt.Println(c1.Frontends[0].Name)

}