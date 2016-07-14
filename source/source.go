package source

import (
	"io/ioutil"
	"os/exec"

	"github.com/pkg/errors"
)

// Source contains all of the details for pulling a repository
type Source struct {
	repo      string
	targetDir string
}

// NewSource instantiates a new source
func NewSource(repo string) *Source {
	return &Source{repo: repo}
}

// Pull pulls the source from git
func (s *Source) Pull() error {
	d, err := ioutil.TempDir("gangway", "pull")
	if err != nil {
		return errors.Wrap(err, "could not create temporary directory for source pull")
	}

	s.targetDir = d
	if _, err := exec.Command("git", "pull", s.repo, s.targetDir).CombinedOutput(); err != nil {
		return errors.Wrap(err, "could not pull source repo")
	}

	return nil
}

// Directory returns the directory that the source was pulled into
func (s *Source) Directory() string {
	return s.targetDir
}
