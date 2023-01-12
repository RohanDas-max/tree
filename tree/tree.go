package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (c *Config) MakeResp(path string) (string, error) {
	var r Report
	resp, r, err := c.tree(path, "", "", "", &r)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%v\n%v directories, %v files",
		resp,
		r.DirCount-(r.FileCount),
		r.FileCount,
	), err
}

func (c *Config) tree(path, indent, line, resp string, r *Report) (string, Report, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return resp, *r, fmt.Errorf("could not stat %s: %v", path, err)
	}

	resp = c.buildResp(path, line, resp, fi)

	if !fi.IsDir() {
		r.FileCount++
		return resp, *r, nil
	}

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return resp, *r, fmt.Errorf("could not read dir %s: %v", path, err)
	}

	var names []string
	//ignoring dot files
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			if c.DirOnly && !fi.IsDir() {
				continue
			}
			r.DirCount++
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

		if resp, *r, err = c.tree(filepath.Join(path, name), indent+add, line, resp, r); err != nil {
			return resp, *r, err
		}
	}

	return resp, *r, nil
}
