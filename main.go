package main

import (
	"flag"
	"fmt"
	"log"
	"tree/tree"
)

func main() {
	arg := "."
	f := flag.Bool("f", false, "f is being used for relative path")
	d := flag.Bool("d", false, "d is being used to read dir only")
	l := flag.Int("l", 0, "l is being used to specify the nested levels")
	p := flag.Bool("p", false, "p is being used to get file permissions")
	flag.Parse()

	if len(flag.Args()) > 0 {
		arg = flag.Args()[0]
	}

	c := tree.Config{RelativePath: *f, DirOnly: *d, Depth: *l, Permission: *p}
	resp, err := c.TreeController(arg)
	if err != nil {
		log.Printf("error occurred while: %v\n", err)
	}

	fmt.Println(resp)
}
