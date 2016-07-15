package builder

import "io"

// BuildOpts contains options for building an image before shipping it off
// to a builder instance
type BuildOpts struct {
	Source       string
	ContextDir   string
	OutputStream io.Writer
	Dockerfile   string
	Tag          string
}
