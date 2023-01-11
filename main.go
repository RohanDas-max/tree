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
	flag.Parse()

	if len(flag.Args()) > 0 {
		arg = flag.Args()[0]
	}

	c := tree.Config{RelativePath: *f, DirOnly: *d}
	resp, err := c.MakeResp(arg)
	if err != nil {
		log.Printf("tree %s: %v\n", arg, err)
	}

	fmt.Println(resp)
}
