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
	d := flag.Bool("d", false, "f is being used to read dir only")
	flag.Parse()

	if len(flag.Args()) > 0 {
		arg = flag.Args()[0]
	}

	var r tree.Report
	c := new(tree.Config)
	c.RelativePath = *f
	c.DirOnly = *d

	resp, r, err := c.Tree(arg, "", "", "", r)
	if err != nil {
		log.Printf("tree %s: %v\n", arg, err)
	}

	fmt.Println(
		fmt.Sprintf(
			"%v\n%v directories, %v files",
			resp,
			r.DirCount,
			r.FileCount,
		),
	)

}
