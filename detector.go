package qiskit

import "github.com/buildpacks/libcnb"

type Detector struct {
}

func (d Detector) Detect(ctx libcnb.DetectContext) (libcnb.DetectResult, error) {
	return libcnb.DetectResult{
		Pass: true,
	}, nil
}
