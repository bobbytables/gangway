package docker

import (
	"github.com/bobbytables/gangway/builder"
	"github.com/bobbytables/gangway/source"
	"github.com/pkg/errors"

	"github.com/fsouza/go-dockerclient"
)

// Builder handles building images against Docker
type Builder struct {
	client *docker.Client
}

var _ builder.Builder = &Builder{}

// NewBuilder returns an instantiated builder
func NewBuilder(c *docker.Client) *Builder {
	return &Builder{client: c}
}

// Build builds a Docker image from a definition
func (b *Builder) Build(bo builder.BuildOpts) *builder.Result {
	opts := docker.BuildImageOptions{
		OutputStream: bo.OutputStream,
		Name:         bo.Tag,
		Dockerfile:   bo.Dockerfile,
	}

	src := source.NewSource(bo.Source)
	if err := src.Pull(); err != nil {
		return &builder.Result{Err: err}
	}

	opts.ContextDir = src.Directory()

	err := b.client.BuildImage(opts)
	if err != nil {
		err = errors.Wrap(err, "could not build image")
	}

	return &builder.Result{Err: err}
}
