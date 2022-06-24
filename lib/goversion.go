package lib

import "runtime"

func Version() string {
	return runtime.Version()
}
