package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (c *Config) Tree(path, indent, line, resp string, r Report) (string, Report, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return resp, r, fmt.Errorf("could not stat %s: %v", path, err)
	}

	if c.RelativePath {
		resp += line + path + "/" + fi.Name() + "\n"
		r.DirCount++
	} else {
		resp += line + fi.Name() + "\n"
		r.DirCount++
	}

	if !fi.IsDir() {
		r.FileCount++
		return resp, r, nil
	}

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return resp, r, fmt.Errorf("could not read dir %s: %v", path, err)
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

		if resp, _, err = c.Tree(filepath.Join(path, name), indent+add, line, resp, r); err != nil {
			return resp, r, err
		}
	}

	return resp, r, nil
}
