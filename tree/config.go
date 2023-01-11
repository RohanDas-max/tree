package tree

type Config struct {
	RelativePath bool
	DirOnly      bool
}

type Report struct {
	DirCount  int
	FileCount int
}
