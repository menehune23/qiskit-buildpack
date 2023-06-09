package qiskit

import (
	"github.com/buildpacks/libcnb"
)

type Builder struct {
}

func (b Builder) Build(ctx libcnb.BuildContext) (libcnb.BuildResult, error) {
	return libcnb.NewBuildResult(), nil
}
