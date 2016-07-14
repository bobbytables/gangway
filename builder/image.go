package builder

import "io"

// BuildOpts contains options for building an image before shipping it off
// to a builder instance
type BuildOpts struct {
	ContextDir   string
	OutputStream io.Writer
	Dockerfile   string
	Tag          string
}
