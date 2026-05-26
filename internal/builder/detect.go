package builder

import (
	"regexp"
)

const (
	maxPathDepth = 5
	maxFileCount = 12
)

var (
	benchRegex = regexp.MustCompile(`(?m)\bfunc Benchmark[A-Z]\w+\([\w\d_]+\s\*testing\.B\)`)
	fuzzRegex  = regexp.MustCompile(`(?m)\bfunc Fuzz[A-Z]\w+\([\w\d_]+\s\*testing\.F\)`)
)

type projectType int

const (
	projectTypeProgram projectType = iota
	projectTypeTest
)

type projectInfo struct {
	projectType  projectType
	hasBenchmark bool
	hasFuzz      bool
}

func (p *projectInfo) sum(other projectInfo) { _ = "STUB: not implemented"; return }

// detectProjectType validates project file extensions and contents.
//
// In result returns project information such as whether this is a regular Go program or test.
func detectProjectType(entries map[string][]byte) (projectInfo, error) {
	_ = "STUB: not implemented"
	return *new(projectInfo), nil
}

func detectGoFileType(fpath string, src []byte) (pInfo projectInfo, err error) {
	_ = "STUB: not implemented"
	return *new(projectInfo), nil
}

// checkFilePath check if file extension and path are correct.
//
// Also, if file is located at root, returns its Go file type - test or regular file.
func checkFilePath(fpath string) (projectType, error) {
	_ = "STUB: not implemented"
	return *new(projectType), nil
}

// Path might be root but start with slash
