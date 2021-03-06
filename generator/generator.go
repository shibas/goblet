package generator

import (
	"github.com/deadcheat/goblet"
)

// Entity generator entity
type Entity struct {
	DirMap  map[string][]string
	FileMap map[string]*goblet.File
	Paths   []string
}

// OptionFlagEntity stores option flag values
type OptionFlagEntity struct {
	IncludePatterns []string
	IgnoreDotFiles  bool
	ExcludeEmptyDir bool
}

// UseCase interface
type UseCase interface {
	LoadFiles(paths []string, option OptionFlagEntity) (*Entity, error)
}

// RegexpRepository repository for slice of regexp
type RegexpRepository interface {
	CompilePatterns(patterns []string) error
	MatchAny(path string) bool
}

// PathMatcherRepository respository to judge path is match or not
type PathMatcherRepository interface {
	Prepare(OptionFlagEntity) error
	Match(path string) bool
}
