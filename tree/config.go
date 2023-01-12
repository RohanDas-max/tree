package tree

import "os"

type Config struct {
	RelativePath bool
	DirOnly      bool
	Depth        int
	Permission   bool
}

type Report struct {
	DirCount  int
	FileCount int
}

func (c *Config) buildResp(path, line, resp string, fi os.FileInfo) string {
	if c.RelativePath {
		resp += line + path + "\n"
	} else if c.Permission {
		resp += line + "[" + fi.Mode().String() + "]" + " " + fi.Name() + "\n"
	} else {
		resp += line + fi.Name() + "\n"
	}
	return resp
}
