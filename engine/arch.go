package engine

import "runtime"

// arch
// windows x86 x64 arm arm64
// linux x86 x64 arm arm64

func detectedOS() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
