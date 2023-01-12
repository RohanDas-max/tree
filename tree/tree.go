package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c *Config) TreeController(path string) (string, error) {
	var r Report
	resp, r, err := c.tree(path, "", "", "", &r, 0)
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

func (c *Config) tree(
	path, indent, line, resp string,
	r *Report,
	depthCount int,
) (string, Report, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return resp, *r, fmt.Errorf("could not stat %s: %v", path, err)
	}

	resp = c.buildResp(path, line, resp, fi)

	if !fi.IsDir() {
		r.FileCount++
		return resp, *r, nil
	}

	if c.Depth != 0 && c.Depth == depthCount {
		return resp, *r, nil
	}

	fis, err := os.ReadDir(path)
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

		if resp, *r, err = c.tree(filepath.Join(path, name), indent+add, line, resp, r, depthCount+1); err != nil {
			return resp, *r, err
		}
	}

	return resp, *r, nil
}
