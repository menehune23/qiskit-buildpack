package main

import (
	"github.com/buildpacks/libcnb"

	qiskit "qiskit-buildpack"
)

func main() {
	libcnb.Main(qiskit.Detector{}, qiskit.Builder{})
}
