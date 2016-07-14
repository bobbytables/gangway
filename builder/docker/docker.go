package docker

import (
	"github.com/bobbytables/gangway/builder"

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
	opts := client.BuildImageOptions{
		OutputStream: bo.OutputStream,
		ContextDir:   bo.ContextDir,
		Name:         bo.Tag,
	}

	return nil
}
