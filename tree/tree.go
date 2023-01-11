package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//This will be the doing the core functionality
func (c *Config) Tree(
	root, indent, line, resp string, r Report) (string, Report, error) {
	fi, err := os.Stat(root)
	if err != nil {
		return resp, r, fmt.Errorf("could not stat %s: %v", root, err)
	}

	if c.RelativePath {
		resp += line + root + "/" + fi.Name() + "\n"
		r.DirCount++
	} else {
		resp += line + fi.Name() + "\n"
		r.DirCount++
	}

	if !fi.IsDir() {
		r.FileCount++
		return resp, r, nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return resp, r, fmt.Errorf("could not read dir %s: %v", root, err)
	}

	var names []string
	//ignoring dot files
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			if c.DirOnly && !fi.IsDir() {
				continue
			}
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := verLine + "  "
		if i == len(names)-1 {
			line = indent + EndLine
			add = "   "
		} else {
			line = indent + vhLine
		}

		if resp, _, err = c.Tree(filepath.Join(root, name), indent+add, line, resp, r); err != nil {
			return resp, r, err
		}
	}

	return resp, r, nil
}
