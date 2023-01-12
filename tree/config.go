package tree

import "os"

type Config struct {
	RelativePath bool
	DirOnly      bool
	Levels       int
}

type Report struct {
	DirCount  int
	FileCount int
}

func (c *Config) buildResp(path, line, resp string, fi os.FileInfo) string {
	if c.RelativePath {
		resp += line + path + "\n"
	} else {
		resp += line + fi.Name() + "\n"
	}
	return resp
}
