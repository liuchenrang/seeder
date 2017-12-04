package bootstrap

import (
	"fmt"
)

type Strapper interface {
	Bootstrap()
}

type LogBootStrapper struct {
	path string
}

func (s *LogBootStrapper) Bootstrap() {
	fmt.Printf("loading log config %v", s.path)
}

func NewLogBootstrapper(path string) *LogBootStrapper {
	return &LogBootStrapper{path: path}
}
